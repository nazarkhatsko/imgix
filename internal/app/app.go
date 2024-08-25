package app

import (
	"encoding/json"
	"fmt"
	"imgix/internal/config"
	"imgix/internal/logger"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/favicon"

	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func InitApp(appConfig config.AppConfig) *fiber.App {
	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       appConfig.Name,
		JSONEncoder:   json.Marshal,
		JSONDecoder:   json.Unmarshal,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"success": false,
				"message": err.Error(),
			})
		},
	})

	if appConfig.Env == config.APP_ENV_DEV {
		app.Get("/metrics", monitor.New())

		app.Use(swagger.New(swagger.Config{
			BasePath: "/api",
			FilePath: "./docs/swagger.yaml",
			Path:     "docs",
		}))
	}

	app.Use(cors.New())

	app.Use(helmet.New())

	app.Use(healthcheck.New())

	app.Use(etag.New())

	app.Use(requestid.New())

	app.Use(logger.Logger(appConfig.Env))

	app.Use(limiter.New(limiter.Config{
		Max:               200,
		Expiration:        60 * time.Second,
		LimiterMiddleware: limiter.SlidingWindow{},
	}))

	app.Get("", func(c *fiber.Ctx) error {
		return c.Redirect("welcome")
	})

	app.Static("/", "./static")

	app.Use(favicon.New(favicon.Config{
		File: "./static/assets/icons/favicon.ico",
		URL:  "/assets/icons/favicon.ico",
	}))

	return app
}

func RunApp(appConfig config.AppConfig, app *fiber.App) error {
	return app.Listen(fmt.Sprintf(":%d", appConfig.Port))
}
