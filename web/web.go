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

		createResult, err := wc.Serv.Create(req)

		if err != nil {
			return err
		}

		return c.JSON(&service.UniversalResponse{
			Response: createResult.Response,
			Error:    createResult.Error,
		})

	})

	wc.App.Get("/delete", func(c *fiber.Ctx) error {
		var req service.DeleteOrderRequest
		if err := c.BodyParser(&req); err != nil {
			return err
		}

		deleteResult, err := wc.Serv.Delete(req.OrderID)

		if err != nil {
			return err
		}

		return c.JSON(&service.UniversalResponse{
			Response: deleteResult.Response,
			Error:    err,
		})

	})
}
