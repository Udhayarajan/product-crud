package mock

import (
	"product-crud/domain"
	"product-crud/model"
	"time"
)

// ErrMock is of int type
type ErrMock int

const (
	DBOperationError ErrMock = iota
	DBDuplicateEntry
	DBNoEntry
	DBUpdateError
	DBDeleteError
	DBAffectedRowMoreThanOneError
)

// ServiceMock ...
type ServiceMock struct {
	Err ErrMock
}

func (s *ServiceMock) InsertProduct(product model.Product) error {
	if s.Err == DBOperationError {
		return domain.ErrProductInsert
	}
	if s.Err == DBDuplicateEntry {
		return domain.ErrProductAlreadyExists
	}
	return nil
}

func (s *ServiceMock) GetProduct(productId string) (*model.ProductResponse, error) {
	if s.Err == DBOperationError {
		return nil, domain.ErrProductGet
	}

	if s.Err == DBNoEntry {
		return nil, domain.ErrProductNotFound
	}
	return &model.ProductResponse{Id: productId}, nil
}

func (s *ServiceMock) GetProducts() (*[]model.ProductResponse, error) {
	if s.Err == DBOperationError {
		return nil, domain.ErrProductsGet
	}
	res := []model.ProductResponse{
		{
			Id:           "",
			CreatedAtUTC: time.Time{},
			UpdatedAtUTC: time.Time{},
			Product:      model.Product{},
		},
	}
	return &res, nil
}

func (s *ServiceMock) UpdateProduct(productId string, product model.Product) error {
	if s.Err == DBUpdateError {
		return domain.ErrProductUpdate
	}

	if s.Err == DBAffectedRowMoreThanOneError {
		return domain.ErrProductUnExpectedAffected
	}

	return nil
}

func (s *ServiceMock) DeleteProduct(productId string) error {
	if s.Err == DBDeleteError {
		return domain.ErrProductDelete
	}
	if s.Err == DBAffectedRowMoreThanOneError {
		return domain.ErrProductUnExpectedAffected
	}
	return nil
}
