package main

import (
	"applytocourier/internal/config"
	"applytocourier/internal/db"
	"applytocourier/internal/service"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	dbClient, err := db.NewPostgresConnection(&config.PostgresConnection{
		Host:     "localhost",
		Port:     5432,
		Database: "ApplyToCourier",
		Password: "postgres",
		Username: "postgres",
	})

	if err != nil {
		infoLog.Fatal("INFO\t", "Error to connect Postgres")
	}

	infoLog.Println("INFO\t", "Succesful connection to Postgres")

	as := service.NewApplyService(dbClient, infoLog)
	app := fiber.New(fiber.Config{
		AppName: "Apply order to courier v0.0.1",
	})

	RegisterRouters(app, as)
	infoLog.Fatal(app.Listen(":1200"))
}

func RegisterRouters(app *fiber.App, as *service.ApplyService) {
	app.Get("/create", func(c *fiber.Ctx) error {
		var req service.CreateOrderRequest
		if err := c.BodyParser(&req); err != nil {
			return err
		}

		createResult, err := as.Create(req)

		if err != nil {
			return err
		}

		return c.JSON(&service.UniversalResponse{
			Response: createResult.Response,
			Error:    createResult.Error,
		})

	})

	app.Get("/delete", func(c *fiber.Ctx) error {
		var req service.DeleteOrderRequest
		if err := c.BodyParser(&req); err != nil {
			return err
		}

		deleteResult, err := as.Delete(req.OrderID)

		if err != nil {
			return err
		}

		return c.JSON(&service.UniversalResponse{
			Response: deleteResult.Response,
			Error:    err,
		})

	})
}
