package routes

import (
	"imgix/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRouteResize(app *fiber.App, resizeHandler *handlers.ResizeHandler) error {
	router := app.Group("resize")

	router.Get("", resizeHandler.Index)

	return nil
}
