package models

import (
	"github.com/google/uuid"
)

type Product struct {
	Quantity  uint32    `json:"quantity"`
	ProductId uuid.UUID `json:"productId"`
}

type Customer struct {
	FirstName       string          `json:"firstName"`
	LastName        string          `json:"lastName"`
	EmailAddresse   string          `json:"emailAddresse"`
	ShippingAddress ShippingAddress `json:"shippingAddress"`
}

type ShippingAddress struct {
	Line1      string `json:"line1"`
	City       string `json:"city"`
	State      string `json:"state"`
	PostalCode string `json:"postalCode"`
}

type Order struct {
	OrderId         uuid.UUID       `json:"orderId"`
	Products        []Product       `json:"products"`
	Customer        Customer        `json:"customer"`
	ShippingAddress ShippingAddress `json:"shippingAddress"`
}
