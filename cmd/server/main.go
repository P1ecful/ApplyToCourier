package main

import (
	"applytocourier/internal/config"
	"applytocourier/internal/db"
	"applytocourier/internal/service"
	"applytocourier/internal/web"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

// Fix get by creator method, add checker for order id, change json strings
func main() {
	logger := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime) // launching the logger

	database := db.NewPostgresConnection(&config.PostgresConnection{
		Host:     "localhost",
		Port:     5432,
		Database: "ApplyToCourier",
		Password: "p1ecful",
		Username: "postgres",
	}, logger) // database connection

	logger.Println("Succesful connection to Postgres")

	wApp := fiber.New()                                          // creating web setup app with fiber
	applyservice := service.NewApplyService(database, logger)    // service setup
	controller := web.CreateNewWebController(wApp, applyservice) // lauching controller setup
	controller.RegisterRouters()                                 // registration routes

	defer database.Close()
	logger.Fatal(wApp.Listen(":1200"))
}
