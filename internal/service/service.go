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

func (s *ApplyService) Create(req CreateOrderRequest) (*UniversalResponse, error) {
	FirstAddress, _ := json.Marshal(req.FirstAddress)
	SecondAddress, _ := json.Marshal(req.SecondAddress)

	sql := `insert into Orders (AuthorID, DeliveryPrice, ItemCategory, 
		ItemWeight, FirstAddressPhone, SecondAddressPhone, CreatedAt, FirstAddress, SecondAddress) 
		values ($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	_, err := s.Database.Exec(sql, req.AuthorID, 300, req.ItemCategory, req.ItemWeight,
		req.FirstAddressPhone, req.SecondAddressPhone, time.Now(), FirstAddress, SecondAddress)

	if err != nil {
		return &UniversalResponse{
			Response: "Ошибка при создании заказа",
			Error:    err,
		}, err
	}

	return &UniversalResponse{
		Response: "Заказ создан",
		Error:    nil,
	}, nil
}

func (s *ApplyService) Delete(OrderID int) (*UniversalResponse, error) {
	sql := `delete from orders where orderid = $1`
	_, err := s.Database.Exec(sql, OrderID)

	if err != nil {
		return &UniversalResponse{
			Response: "Ошибка при удалении заказа",
			Error:    err,
		}, err
	}

	return &UniversalResponse{
		Response: "Ваш заказ удален",
		Error:    nil,
	}, nil
}
