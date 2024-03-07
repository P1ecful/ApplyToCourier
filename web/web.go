package web

import (
	"applytocourier/internal/service"

	"github.com/gofiber/fiber/v2"
)

type WebController struct {
	Serv *service.ApplyService
	App  *fiber.App
}

func CreateNewWebController(app *fiber.App, serv *service.ApplyService) *WebController {
	return &WebController{
		Serv: serv,
		App:  app,
	}
}

// service router
func (wc *WebController) RegisterRouters() {
	wc.App.Get("/create", func(c *fiber.Ctx) error {
		var req service.CreateOrderRequest
		if err := c.BodyParser(&req); err != nil {
			return err
		}

		createResult := wc.Serv.Create(req)

		return c.JSON(&service.UniversalResponse{
			Response: createResult.Response,
		})

	})

	wc.App.Get("/delete", func(c *fiber.Ctx) error {
		var req service.DeleteOrderRequest
		if err := c.BodyParser(&req); err != nil {
			return err
		}

		deleteResult := wc.Serv.Delete(req.OrderID)

		return c.JSON(&service.UniversalResponse{
			Response: deleteResult.Response,
		})

	})
}
