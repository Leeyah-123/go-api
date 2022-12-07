package app

import (
	"log"
	"os"
	"os/signal"

	"fiber/pkg/config"

	"github.com/gofiber/fiber/v2"
)

func StartApp() {
	app := fiber.New(config.FiberConfig())

	// enable middlware
	EnableMiddlewares(app)

	// attach routes
	AttachRoutes(app)

	// Create channel for idle connections.
	idleConnsClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt) // Catch OS signals.
		<-sigint

		// Received an interrupt signal, shutdown.
		if err := app.Shutdown(); err != nil {
			// Error from closing listeners, or context timeout:
			log.Printf("Oops... Server is not shutting down! Reason: %v", err)
		}

		close(idleConnsClosed)
	}()

	// Run server.
	if err := app.Listen(":" + os.Getenv("PORT")); err != nil {
		log.Printf("Oops... Server is not running! Reason: %v", err)
	}

	<-idleConnsClosed
}
