package main

import (
	"applytocourier/internal/config"
	"applytocourier/internal/db"
	"applytocourier/internal/service"
	"applytocourier/internal/web"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

// loads values from .env
func init() {
	if err := godotenv.Load("config/local.env"); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	logger := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime) // launching the logger
	database := db.NewRepository(&config.PostgresConnection{    // database connection
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     os.Getenv("POSTGRES_PORT"),
		Database: os.Getenv("POSTGRES_DB"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Username: os.Getenv("POSTGRES_USER"),
	}, logger).Connect()

	logger.Println("Succesful connection to Postgres")

	wApp := fiber.New()                                                  // creating web setup app with fiber
	applyservice := service.NewApplyService(database, logger)            // service setup
	controller := web.CreateNewWebController(wApp, applyservice, logger) // lauching controller setup
	controller.RegisterRouters()                                         // registration routes

	// start listening and graceful shutdown
	go func() {
		if err := wApp.Listen(":1200"); err != nil {
			logger.Fatal("Error while listening")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	database.Close() // database closing

	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := wApp.Shutdown(); err != nil { // try to stop server
		logger.Print("Failed to stop server")

		return
	}

	logger.Print("Server stopped")
}
