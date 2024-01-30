package service

import (
	"database/sql"
	"encoding/json"
	"log"
	"time"
)

type ApplyToCourierService interface {
	Create(creq CreateOrderRequest) (*UniversalResponse, error)
	Delete(OrderID int) (*UniversalResponse, error)
}

type ApplyService struct {
	Database *sql.DB
	Logger   *log.Logger
}

func NewApplyService(db *sql.DB, lg *log.Logger) *ApplyService {
	return &ApplyService{
		Database: db,
		Logger:   lg,
	}
}

func (s *ApplyService) Create(creq CreateOrderRequest) (*UniversalResponse, error) {
	FirstAddress, _ := json.Marshal(creq.FirstAddress)
	SecondAddress, _ := json.Marshal(creq.SecondAddress)

	sql := `insert into Orders (AuthorID, DeliveryPrice, ItemCategory, 
		ItemWeight, FirstAddressPhone, SecondAddressPhone, CreatedAt, FirstAddress, SecondAddress) 
		values ($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	_, err := s.Database.Exec(sql, creq.AuthorID, 300, creq.ItemCategory, creq.ItemWeight,
		creq.FirstAddressPhone, creq.SecondAddressPhone, time.Now(), FirstAddress, SecondAddress)

	if err != nil {
		return &UniversalResponse{
			Response: "Ошибка при создании заказа",
			Error:    err,
		}, nil
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
		}, nil
	}

	return &UniversalResponse{
		Response: "Ваш заказ удален",
		Error:    nil,
	}, nil
}
