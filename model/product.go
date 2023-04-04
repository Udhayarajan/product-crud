package model

import "time"

type Product struct {
	Name        string      `json:"name" validate:"required,min=5,max=100"`
	Type        ProductType `json:"type" validate:"required,oneof=TypeA TypeB TypeC TypeD"`
	Quantity    int         `json:"quantity" validate:"min=0"`
	Price       float64     `json:"price" validate:"min=0"`
	Description *string     `json:"description" validate:"omitempty,max=500"`
}

type ProductResponse struct {
	Id           string    `json:"id" validate:"uuid"`
	CreatedAtUTC time.Time `json:"created_at_utc"`
	UpdatedAtUTC time.Time `json:"updated_at_utc"`
	Product
}
