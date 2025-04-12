package models

type Product struct {
	Quantity  uint32 `json:"quantity"`
	ProductId string `json:"productId"`
}

type Customer struct {
	FistName        string          `json:"firstName"`
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
