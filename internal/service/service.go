package service

import (
	"database/sql"
	"encoding/json"
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

func (s *ApplyService) Create(req CreateOrderRequest) *UniversalResponse {
	//var id int
	sql := `insert into Orders (AuthorID, DeliveryPrice, ItemCategory, 
		ItemWeight, FirstAddressPhone, SecondAddressPhone, CreatedAt, FirstAddress, SecondAddress) 
		values ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id`

	FirstAddress, _ := json.Marshal(req.FirstAddress)
	SecondAddress, _ := json.Marshal(req.SecondAddress)

	// if err := s.Database.QueryRow(sql, req.AuthorID, 300, req.ItemCategory, req.ItemWeight,
	// 	req.FirstAddressPhone, req.SecondAddressPhone, time.Now(),
	// 	FirstAddress, SecondAddress).Scan(&id); err != nil {
	// 	return &UniversalResponse{
	// 		Response: "Error with creating order",
	// 	}
	// }

	_, err := s.Database.Exec(sql, req.AuthorID, 300, req.ItemCategory, req.ItemWeight,
		req.FirstAddressPhone, req.SecondAddressPhone, time.Now(), FirstAddress, SecondAddress)

	if err != nil {
		s.Logger.Println(err)

		return &UniversalResponse{
			Response: "Error with creating order",
		}
	}

	return &UniversalResponse{
		Response: "Order created",
	}
}

func (s *ApplyService) Delete(OrderID int) *UniversalResponse {
	sql := `delete from orders where orderid = $1`
	_, err := s.Database.Exec(sql, OrderID)

	if err != nil {
		return &UniversalResponse{
			Response: "Error with deleting order",
		}
	}

	return &UniversalResponse{
		Response: "Order was deleted",
	}
}
