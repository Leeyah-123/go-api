package router

import (
	"fiber/app/handler"
	"fiber/pkg/middleware"

	"github.com/gofiber/fiber/v2"
)

func TodoRouter(app fiber.Router) {
	authRoute := app.Group("/todos")

	authRoute.Get("/", middleware.JWTProtected(), handler.GetUserTodos)
	authRoute.Get("/:id", middleware.JWTProtected(), handler.GetTodoById)
	authRoute.Post("/", middleware.JWTProtected(), handler.AddTodo)
	authRoute.Patch("/:id", middleware.JWTProtected(), handler.UpdateTodoById)
	authRoute.Delete("/:id", middleware.JWTProtected(), handler.DeleteTodoById)
}
