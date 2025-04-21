package models

import (
	"github.com/google/uuid"
)

type Product struct {
	Quantity    uint32 `json:"quantity" binding:"required"`
	ProductCode string `json:"productCode" binding:"required"`
}

type Customer struct {
	FirstName       string          `json:"firstName" binding:"required"`
	LastName        string          `json:"lastName" binding:"required"`
	EmailAddress    string          `json:"emailAddress" binding:"required,email"`
	ShippingAddress ShippingAddress `json:"shippingAddress" binding:"required"`
}

type ShippingAddress struct {
	Line1      string `json:"line1" binding:"required"`
	City       string `json:"city" binding:"required"`
	State      string `json:"state" binding:"required"`
	PostalCode string `json:"postalCode" binding:"required"`
}

type Order struct {
	OrderId  uuid.UUID `json:"orderId" binding:"required"`
	Products []Product `json:"products" binding:"required"`
	Customer Customer  `json:"customer" binding:"required"`
}
