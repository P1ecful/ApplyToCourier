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
	wc.app.Get("/create", func(c *fiber.Ctx) error {
		var req service.CreateOrUpdateRequest
		if err := c.BodyParser(&req); err != nil {
			wc.log.Fatal(err)
		}

		return c.JSON(wc.serv.Create(req))

	})

	wc.app.Get("/delete", func(c *fiber.Ctx) error {
		var req service.DeleteRequest
		if err := c.BodyParser(&req); err != nil {
			wc.log.Fatal(err)
		}

		return c.JSON(wc.serv.Delete(req.OrderID))
	})

	wc.app.Get("/get-creator", func(c *fiber.Ctx) error {
		var req service.GetByCreatorRequest
		if err := c.BodyParser(&req); err != nil {
			wc.log.Fatal(err)
		}

		return c.JSON(wc.serv.GetWithCreatorID(req.CreatorID))
	})
}
