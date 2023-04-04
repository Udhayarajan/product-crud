package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"product-crud/handler"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc gin.HandlerFunc
}

type Routes []Route

func NewRoutes(ph handler.ProductHandler) Routes {
	return Routes{
		Route{
			Name:        "Insert Product",
			Method:      http.MethodPost,
			Pattern:     "/products",
			HandlerFunc: ph.InsertProductHandler,
		},
		Route{
			Name:        "Get single product",
			Method:      http.MethodGet,
			Pattern:     "/products/:id",
			HandlerFunc: ph.GetProductHandler,
		},
		Route{
			Name:        "Get All products",
			Method:      http.MethodGet,
			Pattern:     "/products",
			HandlerFunc: ph.GetProductsHandler,
		},
		Route{
			Name:        "Update product",
			Method:      http.MethodPut,
			Pattern:     "/products/:id",
			HandlerFunc: ph.UpdateProductHandler,
		},
		Route{
			Name:        "Delete product",
			Method:      http.MethodDelete,
			Pattern:     "/products/:id",
			HandlerFunc: ph.DeleteProductHandler,
		},
	}
}
func AttachRoutes(server *gin.Engine, routes Routes) {
	for _, route := range routes {
		server.Handle(route.Method, route.Pattern, route.HandlerFunc)
	}
}
