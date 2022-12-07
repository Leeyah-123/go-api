package config

import (
	"errors"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

func FiberConfig() fiber.Config {
	return fiber.Config{
		UnescapePath: true,
		JSONEncoder:  json.Marshal,
		JSONDecoder:  json.Unmarshal,
		// Override default error handler
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			// Status code defaults to 500
			code := fiber.StatusInternalServerError

			// Retrieve the custom status code if it's a *fiber.Error
			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}

			// Send custom error
			err = ctx.Status(code).JSON(fiber.Map{
				"message": "An error occurred: " + e.Message,
			})
			if err != nil {
				// In case the SendFile fails
				return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
			}

			// Return from handler
			return nil
		},
	}
}
