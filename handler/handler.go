package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"product-crud/domain"
	"reflect"
	"strings"
)

var (
	validate = validator.New()
)

func init() {
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
}

type Handler interface {
	InsertProductHandler(c *gin.Context)
	GetProductHandler(c *gin.Context)
	GetProductsHandler(c *gin.Context)
	UpdateProductHandler(c *gin.Context)
	DeleteProductHandler(c *gin.Context)
}

type ProductHandler struct {
	domain domain.Service
}

func NewProductHandler(service domain.Service) ProductHandler {
	return ProductHandler{
		domain: service,
	}
}
