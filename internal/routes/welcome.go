package routes

import (
	"imgix/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRouteWelcome(app *fiber.App, welcomeHandler *handlers.WelcomeHandler) error {
	router := app.Group("welcome")

	router.Get("", welcomeHandler.Index)

	return nil
}
