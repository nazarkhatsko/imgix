package routes

import (
	"imgix/internal/middlewares"

	"github.com/gofiber/fiber/v2"
)

func SetupRouteNotFound(app *fiber.App) error {
	app.Use(middlewares.NotFoundMiddleware)

	return nil
}
