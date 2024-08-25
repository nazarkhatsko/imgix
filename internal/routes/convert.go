package routes

import (
	"imgix/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRouteConvert(app *fiber.App, convertHandler *handlers.ConvertHandler) error {
	router := app.Group("convert")

	router.Get("", convertHandler.Index)

	return nil
}
