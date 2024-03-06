package main

import (
	"applytocourier/internal/config"
	"applytocourier/internal/db"
	"applytocourier/internal/logging"
	"applytocourier/internal/service"
	"applytocourier/web"

	"github.com/gofiber/fiber/v2"
)

// TODO: migrations, order update, search order by author_ID
func main() {
	logger := logging.NewLogger() // launching the logger

	database, err := db.NewPostgresConnection(&config.PostgresConnection{
		Host:     "localhost",
		Port:     5432,
		Database: "ApplyToCourier",
		Password: "postgres",
		Username: "postgres",
	}) // database connection

	if err != nil {
		logger.Fatal("INFO\t", "Error to connect Postgres")
	}

	logger.Println("INFO\t", "Succesful connection to Postgres")
	//db.MakeMigrations() // migrations for postgres !FIXME

	wApp := fiber.New()                                          // creating web setup app with fiber
	applyservice := service.NewApplyService(database, logger)    // service setup
	controller := web.CreateNewWebController(wApp, applyservice) // lauching controller setup
	controller.RegisterRouters()                                 // registration routes

	logger.Fatal(wApp.Listen(":1200"))
}
