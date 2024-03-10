package service

import (
	"applytocourier/internal/db/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type ApplyService struct {
	Database *sql.DB
	Logger   *log.Logger
}

func NewApplyService(db *sql.DB, log *log.Logger) *ApplyService {
	return &ApplyService{
		Database: db,
		Logger:   log,
	}
}

type CreateOrUpdateRequest struct {
	AuthorID           int            `json:"author id"`
	ItemCategory       string         `json:"item category"`
	ItemWeight         string         `json:"item weight"`
	FirstAddressPhone  string         `json:"first address phone"`
	SecondAddressPhone string         `json:"second address phone"`
	FirstAddress       models.Address `json:"first address"`
	SecondAddress      models.Address `json:"second address"`
}

type DeleteRequest struct {
	OrderID int `json:"order id"`
}

type UniversalResponse struct {
	Response string `json:"response"`
}

func (s *ApplyService) Create(req CreateOrUpdateRequest) *UniversalResponse {
	var id int

	FirstAddress, _ := json.Marshal(req.FirstAddress)
	SecondAddress, _ := json.Marshal(req.SecondAddress)

	sql := `insert into orders (AuthorID, DeliveryPrice, ItemCategory, 
		ItemWeight, FirstAddressPhone, SecondAddressPhone, CreatedAt, FirstAddress, SecondAddress) 
		values ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING orderid`

	if err := s.Database.QueryRow(sql, req.AuthorID, 300, req.ItemCategory, req.ItemWeight, req.FirstAddressPhone, req.SecondAddressPhone, time.Now(), FirstAddress, SecondAddress).Scan(&id); err != nil {
		return &UniversalResponse{
			Response: "Error with creating order",
		}
	}

	return &UniversalResponse{
		Response: fmt.Sprintf("Order %d created", id),
	}
}

func (s *ApplyService) Delete(OrderId int) *UniversalResponse {
	sql := `delete from orders where orderid = $1`
	_, err := s.Database.Exec(sql, OrderId)

	if err != nil {
		return &UniversalResponse{
			Response: "Error with deleting order",
		}
	}

	return &UniversalResponse{
		Response: "Order was deleted",
	}
}
