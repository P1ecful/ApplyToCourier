package service

import "applytocourier/internal/db/models"

type CreateOrderRequest struct {
	AuthorID           int            `json:"author id"`
	ItemCategory       string         `json:"item category"`
	ItemWeight         string         `json:"item weight"`
	FirstAddressPhone  string         `json:"first address phone"`
	SecondAddressPhone string         `json:"second address phone"`
	FirstAddress       models.Address `json:"first address"`
	SecondAddress      models.Address `json:"second address"`
}

type DeleteOrderRequest struct {
	OrderID int `json:"order id"`
}

type UpdateOrderRequest struct {
	OrderID int `json:"order id"`
}

type UniversalResponse struct {
	Response string `json:"response"`
}
