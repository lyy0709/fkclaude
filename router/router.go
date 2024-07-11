package router

import (
	"fkclaude/serve"
	"fkclaude/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	app.Use(logger.New(logger.ConfigDefault))

	// Serve the login page
	app.Get("/login", handlers.LoginPage)
	// Handle session key submission
	app.Post("/login", handlers.Login)

	// Handle all other requests
	serve.APIHandler(app)
}

