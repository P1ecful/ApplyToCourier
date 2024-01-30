package models

import "time"

type Order struct {
	AuthorID           int       `json:"author id"`
	OrderID            int       `json:"order id"`
	DeliveryPrice      int       `json:"price"`
	ItemCategory       string    `json:"item category"`
	ItemWeight         string    `json:"item weight"`
	FirstAddressPhone  string    `json:"first address phone"`
	SecondAddressPhone string    `json:"second address phone"`
	FirstAddress       Address   `json:"first address"`
	SecondAddress      Address   `json:"second address"`
	CreatedAt          time.Time `json:"createdAt"`
}

type Address struct {
	Street       string `json:"street"`
	Home         int    `json:"home"`
	Housing      int    `json:"housing"`
	Entrance     int    `json:"enrance"`
	Floor        int    `json:"floor"`
	Flat         int    `json:"flat"`
	IntercomCode string `json:"intercom code"`
}
