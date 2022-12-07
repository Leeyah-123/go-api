package router

import (
	"fiber/app/handler"
	"fiber/pkg/middleware"

	"github.com/gofiber/fiber/v2"
)

func AuthenticationRouter(app fiber.Router) {
	authRoute := app.Group("/auth")

	authRoute.Get("/profile", middleware.JWTProtected(), handler.Profile)
	authRoute.Post("/login", handler.Login)
	authRoute.Post("/signup", handler.SignUp)
}
