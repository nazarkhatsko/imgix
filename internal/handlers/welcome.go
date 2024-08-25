package handlers

import (
	"imgix/internal/common/constants"
	"imgix/internal/common/utils"
	"imgix/internal/components/pages"
	"imgix/internal/config"
	"imgix/internal/render"

	"github.com/gofiber/fiber/v2"
)

type WelcomeHandler struct {
	appConfig config.AppConfig
}

func NewWelcomeHandler(appConfig config.AppConfig) *WelcomeHandler {
	handler := WelcomeHandler{
		appConfig: appConfig,
	}
	return &handler
}

func (handler *WelcomeHandler) Index(c *fiber.Ctx) error {
	options := pages.WelcomeOptions{
		AppName:    handler.appConfig.Name,
		AppVersion: handler.appConfig.Version,
		PageName:   constants.WELLCOME_PAGE_NAME,
	}

	return render.Render(c, utils.GetFormatPageTitle(handler.appConfig.Name, options.PageName), pages.Welcome(options))
}
