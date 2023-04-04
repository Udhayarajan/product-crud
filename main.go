package main

import (
	"database/sql"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"os/signal"
	"product-crud/domain"
	"product-crud/handler"
	"product-crud/routes"
	"time"
)

var (
	buildType        string // debug/release
	buildVersion     string
	buildVersionCode int
	buildTime        string

	helpFlag    bool
	versionFlag bool

	port string

	errChan = make(chan error)
	done    = make(chan bool)
)

func init() {
	flag.BoolVar(&versionFlag, "version", false, "show current version and exit")
	flag.BoolVar(&helpFlag, "help", false, "show usage and exit")
	flag.StringVar(&port, "port", ":4444", "server port")
}

func main() {
	setBuildVariable()
	parseFlag()
	go handleInterrupts()

	ginServer := gin.Default()
	db, err := openDB()
	if err != nil {
		log.Printf("error connecting DB: %v", err)
		return
	}
	log.Println("DB connection is successful")
	defer db.Close()

	productService := domain.NewProductService(db)
	productHandler := handler.NewProductHandler(&productService)

	apiRoutes := routes.NewRoutes(productHandler)
	routes.AttachRoutes(ginServer, apiRoutes)

	go func() {
		errChan <- ginServer.Run(port)
	}()
	select {
	case err := <-errChan:
		log.Printf("ListenAndServe error: %v", err)
	case <-done:
		log.Println("shutting down server ...")
	}

	time.AfterFunc(1*time.Second, func() {
		close(done)
		close(errChan)
	})
}

func openDB() (*sql.DB, error) {
	var (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "root"
		dbname   = "product-management"
	)

	connectionString := os.Getenv("POSTGRESQL_CONN_STRING") // for production build
	if len(connectionString) == 0 {
		connectionString = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	}
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func handleInterrupts() {
	log.Println("start handle interrupts")
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
}

func parseFlag() {
	flag.Parse()
	if helpFlag {
		flag.Usage()
		os.Exit(0)
	}

	if versionFlag {
		fmt.Printf("%s %s(%d) %s\n", buildType, buildVersion, buildVersionCode, buildTime)
		os.Exit(0)
	}
}

func setBuildVariable() {
	if buildType == "" {
		buildType = "debug"
	}

	if buildVersion == "" {
		buildVersion = "0.0.1"
	}
	if buildVersionCode == 0 {
		buildVersionCode = 1
	}
	if buildTime == "" {
		buildTime = time.Now().UTC().Format(time.RFC3339)
	}
}
