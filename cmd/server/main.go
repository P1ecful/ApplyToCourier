package main

import (
	"applytocourier/internal/config"
	"applytocourier/internal/db"
	"applytocourier/internal/service"
	"applytocourier/web"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

// TODO: order update, search order by author_ID, returning id with response
func main() {
	logger := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime) // launching the logger

	database, err := db.NewPostgresConnection(&config.PostgresConnection{
		Host:     "localhost",
		Port:     5432,
		Database: "ApplyToCourier",
		Password: "postgres",
		Username: "postgres",
	}) // database connection

	if err != nil {
		logger.Fatal("Error to connect Postgres")
	}

	logger.Println("Succesful connection to Postgres")

	wApp := fiber.New()                                          // creating web setup app with fiber
	applyservice := service.NewApplyService(database, logger)    // service setup
	controller := web.CreateNewWebController(wApp, applyservice) // lauching controller setup
	controller.RegisterRouters()                                 // registration routes

	defer database.Close()
	logger.Fatal(wApp.Listen(":1200"))
}
