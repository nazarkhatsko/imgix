package logger

import (
	"imgix/internal/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Logger(appEnv string) func(*fiber.Ctx) error {
	switch appEnv {
	case config.APP_ENV_PROD:
		return logger.New(logger.Config{
			DisableColors: true,
			// JSONDecoder:
			Format: "${pid} ${locals:requestid} ${status} - ${method} ${path}\n",
		})
	case config.APP_ENV_DEV:
		return logger.New(logger.Config{
			Format: "${pid} ${locals:requestid} ${status} - ${method} ${path}\n",
		})
	case config.APP_ENV_DEBUG:
		return logger.New(logger.Config{
			Format: "${pid} ${locals:requestid} ${status} - ${method} ${path}\n",
		})
	default:
		return logger.New(logger.Config{})
	}
}
