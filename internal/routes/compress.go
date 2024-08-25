package routes

import (
	"imgix/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRouteCompress(app *fiber.App, compressHandler *handlers.CompressHandler) error {
	router := app.Group("compress")

	router.Get("", compressHandler.Index)

	return nil
}
