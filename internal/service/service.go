package service

import (
	"applytocourier/internal/db/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
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

// Requests and Responses
type CreateOrUpdateRequest struct {
	CreatorID          int            `json:"creator id"`
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

type GetByCreatorRequest struct {
	CreatorID int `json:"creator id"`
}

type UniversalResponse struct {
	OK       bool   `json:"OK"`
	Response string `json:"Response"`
	Method   string `json:"Method"`
	Error    error  `json:"Error"`
}

// Services
func (s *ApplyService) Create(req CreateOrUpdateRequest) *UniversalResponse {
	generatedID := rand.Intn(999999-100000) + 100000 // generating ID for order

	// Marshalization addresses to JSON format
	FirstAddress, _ := json.Marshal(req.FirstAddress)
	SecondAddress, _ := json.Marshal(req.SecondAddress)

	sql := `insert into orders (OrderID, CreatorID, DeliveryPrice, ItemCategory, 
		ItemWeight, FirstAddressPhone, SecondAddressPhone, CreatedAt, FirstAddress, SecondAddress) 
		values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`

	_, err := s.Database.Exec(sql, generatedID, req.CreatorID, 300, req.ItemCategory, req.ItemWeight, req.FirstAddressPhone, req.SecondAddressPhone, time.Now(), FirstAddress, SecondAddress)

	if err != nil { // error response
		return &UniversalResponse{
			OK:       false,
			Response: "Error with creating order",
			Method:   "create",
			Error:    err,
		}
	}

	return &UniversalResponse{
		OK:       true,
		Response: fmt.Sprintf("Order %d created", generatedID),
		Method:   "create",
		Error:    err,
	}
}

func (s *ApplyService) Delete(orderid int) *UniversalResponse {
	_, err := s.Database.Exec(`delete from orders where orderid = $1`, orderid)

	if err != nil {
		return &UniversalResponse{ // error response
			OK:       false,
			Response: fmt.Sprintf("Order %d wasn't deleted", orderid),
			Method:   "delete",
			Error:    err,
		}
	}

	return &UniversalResponse{
		OK:       true,
		Response: fmt.Sprintf("Order %d was deleted", orderid),
		Method:   "delete",
		Error:    err,
	}
}

func (s *ApplyService) GetByCreatorID(creatorid int) *models.Order { // !FIXME
	var order *models.Order
	sql := `select * from orders where creatorid = $1`

	if err := s.Database.QueryRow(sql, creatorid).Scan(&order); err != nil {
		s.Logger.Fatal(err)
	}

	return order
}
