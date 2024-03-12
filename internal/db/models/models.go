package models

import "time"

type Order struct {
	CreatorID          int       `json:"Creator Id"`
	OrderID            int       `json:"Order Id"`
	DeliveryPrice      int       `json:"Price"`
	ItemCategory       string    `json:"Item Category"`
	ItemWeight         string    `json:"Item Weight"`
	FirstAddressPhone  string    `json:"First Address Phone"`
	SecondAddressPhone string    `json:"Second Address Phone"`
	FirstAddress       Address   `json:"First Address"`
	SecondAddress      Address   `json:"Second Address"`
	CreatedAt          time.Time `json:"Created At"`
}

type Address struct {
	Street       string `json:"Street"`
	Home         int    `json:"Home"`
	Housing      int    `json:"Housing"`
	Entrance     int    `json:"Enrance"`
	Floor        int    `json:"Floor"`
	Flat         int    `json:"Flat"`
	IntercomCode string `json:"Intercom Code"`
}
