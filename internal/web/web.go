package web

import (
	"applytocourier/internal/service"
	"log"

	"github.com/gofiber/fiber/v2"
)

type WebController struct {
	serv *service.ApplyService
	app  *fiber.App
	log  *log.Logger
}

func CreateNewWebController(app *fiber.App, serv *service.ApplyService, log *log.Logger) *WebController {
	return &WebController{
		serv: serv,
		app:  app,
		log:  log,
	}
}

// service's handlers
func (wc *WebController) RegisterRouters() {
	// http://127.0.0.1:1200/create
	wc.app.Post("/create", func(c *fiber.Ctx) error {
		var req service.CreateOrUpdateRequest
		if err := c.BodyParser(&req); err != nil {
			wc.log.Fatal(err)
		}

		return c.JSON(wc.serv.Create(req))

	})

	// http://127.0.0.1:1200/delete
	wc.app.Post("/delete", func(c *fiber.Ctx) error {
		var req service.OrderRequest
		if err := c.BodyParser(&req); err != nil {
			wc.log.Fatal(err)
		}

		return c.JSON(wc.serv.Delete(req.OrderID))
	})

	// http://127.0.0.1:1200/get-creator
	wc.app.Post("/get-creator", func(c *fiber.Ctx) error {
		var req service.GetCreatorRequest
		if err := c.BodyParser(&req); err != nil {
			wc.log.Fatal(err)
		}

		return c.JSON(wc.serv.GetWithCreatorID(req.CreatorID))
	})

	// http://127.0.0.1:1200/get-order
	wc.app.Post("/get-order", func(c *fiber.Ctx) error {
		var req service.OrderRequest

		if err := c.BodyParser(&req); err != nil {
			wc.log.Fatal(err)
		}

		return c.JSON(wc.serv.GetWithOrderID(req.OrderID))
	})

}
