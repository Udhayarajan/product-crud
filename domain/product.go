package domain

import (
	"database/sql"
	"errors"
	"github.com/lib/pq"
	"log"
	"product-crud/model"
	"time"
)

var (
	ErrProductAlreadyExists      = errors.New("product already exists")
	ErrProductNotFound           = errors.New("requested product not found")
	ErrProductInsert             = errors.New("failed to create the product")
	ErrProductNameTooShort       = errors.New("product name must be greater that 5 character")
	ErrProductPriceNegative      = errors.New("product price must be greater or equal to 0")
	ErrProductGet                = errors.New("failed to get the product")
	ErrProductsGet               = errors.New("unknown error occurred while getting products")
	ErrProductUpdate             = errors.New("failed to update the product")
	ErrProductUnExpectedAffected = errors.New("unexpected number of rows modified")
	ErrProductDelete             = errors.New("failed to delete the product")
)

const (
	postgresUniqueConstraintViolationCode = "23505"
	postgresNoDataFoundErrorCode          = "P0002"
)

func (ps *ProductService) InsertProduct(product model.Product) error {
	stmt := "INSERT INTO products (name, type, quantity, price, description) VALUES ($1, $2, $3, $4, $5)"
	_, err := ps.db.Exec(stmt, product.Name, product.Type, product.Quantity, product.Price, product.Description)
	if err != nil {
		pqErr, ok := err.(*pq.Error)
		if ok {
			if pqErr.Code == postgresUniqueConstraintViolationCode {
				switch pqErr.Constraint {
				case "products_pkey":
					return ErrProductAlreadyExists
				default:
					log.Println("[ERROR:InsertProduct]:", pqErr.Message)
					return ErrProductInsert
				}
			}
			return ErrProductInsert
		}
		return err
	}
	return nil
}

func (ps *ProductService) GetProduct(productId string) (*model.ProductResponse, error) {
	product := model.ProductResponse{}
	err := ps.db.QueryRow("SELECT id,name,quantity,type, price, description, createdat, updatedat FROM products WHERE id=$1", productId).Scan(&product.Id, &product.Name, &product.Quantity, &product.Type, &product.Price, &product.Description, &product.CreatedAtUTC, &product.UpdatedAtUTC)
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return nil, ErrProductNotFound
		}
		log.Println("[ERROR:GetProduct]", err)
		return nil, ErrProductGet
	}
	return &product, nil
}

func (ps *ProductService) GetProducts() (*[]model.ProductResponse, error) {
	rows, err := ps.db.Query("SELECT id,name,quantity,type, price, description, createdat, updatedat FROM products")
	if err != nil {
		log.Println("[ERROR: GetProducts]", err)
		return nil, ErrProductsGet
	}
	var products []model.ProductResponse
	for rows.Next() {
		product := model.ProductResponse{}
		err := rows.Scan(&product.Id, &product.Name, &product.Quantity, &product.Type, &product.Price, &product.Description, &product.CreatedAtUTC, &product.UpdatedAtUTC)
		if err != nil {
			log.Println("[ERROR: GetProducts]", err)
			return nil, ErrProductsGet
		}
		products = append(products, product)
	}
	return &products, nil
}
func (ps *ProductService) UpdateProduct(productId string, product model.Product) error {
	stmt := "UPDATE products SET name=$1, quantity=$2, type=$3, price=$4, description=$5, updatedat=$7 WHERE id=$6"
	exec, err := ps.db.Exec(stmt, product.Name, product.Quantity, product.Type, product.Price, product.Description, productId, time.Now().UTC())
	if err != nil {
		pqErr, ok := err.(*pq.Error)
		if ok {
			log.Println("[ERROR:UpdateProduct]", "(pq.Error)", pqErr)
			return ErrProductUpdate
		}
		log.Println("[ERROR:UpdateProduct]", err)
		return ErrProductUpdate
	}
	affected, err := exec.RowsAffected()
	if err != nil {
		log.Println("[ERROR:UpdateProduct]", "RowsAffected", err)
		return ErrProductUpdate
	}
	if affected != 1 {
		return ErrProductUnExpectedAffected
	}
	return nil
}

func (ps *ProductService) DeleteProduct(productId string) error {
	exec, err := ps.db.Exec("DELETE FROM products WHERE id=$1", productId)
	if err != nil {
		pqErr, ok := err.(*pq.Error)
		if ok {
			if pqErr.Code == postgresNoDataFoundErrorCode {
				return ErrProductNotFound
			} else {
				log.Println("[Error: DeleteProduct]", "(pq.Error)", pqErr)
				return ErrProductDelete
			}
		}
		log.Println("[Error: DeleteProduct]", err)
		return ErrProductDelete
	}
	affected, err := exec.RowsAffected()
	if err != nil {
		log.Println("[Error: DeleteProduct]", err)
		return ErrProductDelete
	}
	if affected != 1 {
		return ErrProductUnExpectedAffected
	}
	return nil
}
