package domain

import (
	"database/sql"
	"product-crud/model"
)

type Service interface {
	InsertProduct(product model.Product) error
	GetProduct(prodId string) (*model.ProductResponse, error)
	GetProducts() (*[]model.ProductResponse, error)
	UpdateProduct(prodId string, product model.Product) error
	DeleteProduct(prodId string) error
}

type ProductService struct {
	db *sql.DB
}

func NewProductService(db *sql.DB) ProductService {
	return ProductService{db: db}
}
