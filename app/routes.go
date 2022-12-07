package app

import (
	"fiber/pkg/router"

	"github.com/gofiber/fiber/v2"
)

func AttachRoutes(app *fiber.App) {
	apiRoute := app.Group("/api")

	apiRoute.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Index route",
		})
	})

	// routers
	router.AuthenticationRouter(apiRoute)
	router.TodoRouter(apiRoute)
}
