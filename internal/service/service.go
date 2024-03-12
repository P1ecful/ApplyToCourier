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
	database *sql.DB
	log      *log.Logger
}

func NewApplyService(db *sql.DB, log *log.Logger) *ApplyService {
	return &ApplyService{
		database: db,
		log:      log,
	}
}

// Requests and Responses
type CreateOrUpdateRequest struct {
	CreatorID          int            `json:"Creator Id"`
	ItemCategory       string         `json:"Item Category"`
	ItemWeight         string         `json:"Item Weight"`
	FirstAddressPhone  string         `json:"First Address Phone"`
	SecondAddressPhone string         `json:"Second Address Phone"`
	FirstAddress       models.Address `json:"First Address"`
	SecondAddress      models.Address `json:"Second Address"`
}

type OrderRequest struct {
	OrderID int `json:"Order Id"`
}

type GetCreatorRequest struct {
	CreatorID int `json:"Creator Id"`
}

type Response struct {
	OK       bool   `json:"OK"`
	Method   string `json:"Method"`
	Error    error  `json:"Error"`
	Response string `json:"Response"`
}

// Services methods
func (s *ApplyService) Create(req CreateOrUpdateRequest) *Response {
	generatedID := rand.Intn(999999-100000) + 100000 // generating ID for order

	// Marshalization addresses to JSON format
	FirstAddress, _ := json.Marshal(req.FirstAddress)
	SecondAddress, _ := json.Marshal(req.SecondAddress)

	sql := `insert into orders (OrderID, CreatorID, DeliveryPrice, ItemCategory, 
		ItemWeight, FirstAddressPhone, SecondAddressPhone, CreatedAt, FirstAddress, SecondAddress) 
		values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`

	_, err := s.database.Exec(sql, generatedID, req.CreatorID, 300, req.ItemCategory, req.ItemWeight, req.FirstAddressPhone, req.SecondAddressPhone, time.Now(), FirstAddress, SecondAddress)

	if err != nil { // error response
		s.log.Printf("Error while inserting to db. %s", err)
	}

	return &Response{
		OK:       true,
		Method:   "create",
		Response: fmt.Sprintf("Order %d created", generatedID),
	}
}

func (s *ApplyService) Delete(orderid int) *Response {
	_, err := s.database.Exec(`delete from orders where orderid = $1`, orderid)

	if err != nil {
		s.log.Printf("Error with deleting from db. %s", err)
	}

	return &Response{
		OK:       true,
		Method:   "delete",
		Response: fmt.Sprintf("Order %d was deleted", orderid),
	}
}

func (s *ApplyService) GetWithCreatorID(creatorid int) *Response {
	var id int
	collection := []int{} // array for database records
	sql := `select orderid from orders where creatorid = $1`

	rows, err := s.database.Query(sql, creatorid)

	if err != nil {
		s.log.Printf("Error while selecting by CreatorID from db. %s", err)
	}

	// add the found ID to the array
	for rows.Next() {
		rows.Scan(&id)
		collection = append(collection, id)
	}

	return &Response{
		OK:       true,
		Method:   "get-creator",
		Response: (fmt.Sprint(collection)),
	}
}

func (s *ApplyService) GetWithOrderID(orderid int) *models.Order {
	var order *models.Order = &models.Order{}
	var firstaddress, secondaddress []byte // addresses in json-view
	sql := `select * from orders where orderid = $1`

	if err := s.database.QueryRow(sql, orderid).Scan(&order.OrderID, &order.CreatorID, &order.DeliveryPrice, &order.ItemCategory, &order.ItemWeight, &order.FirstAddressPhone, &order.SecondAddressPhone, &order.CreatedAt, &firstaddress, &secondaddress); err != nil {
		s.log.Printf("Error while scaning order. %s", err)
	}

	// unmarshaling addresses
	json.Unmarshal(firstaddress, &order.FirstAddress)
	json.Unmarshal(secondaddress, &order.SecondAddress)

	return order
}
