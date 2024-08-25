package handlers

import (
	"imgix/internal/common/constants"
	"imgix/internal/common/utils"
	"imgix/internal/components/pages"
	"imgix/internal/config"
	"imgix/internal/render"

	"github.com/gofiber/fiber/v2"
)

type ConvertHandler struct {
	appConfig config.AppConfig
}

func NewConvertHandler(appConfig config.AppConfig) *ConvertHandler {
	handler := ConvertHandler{
		appConfig: appConfig,
	}
	return &handler
}

func (handler *ConvertHandler) Index(c *fiber.Ctx) error {
	options := pages.ConvertOptions{
		AppName:  handler.appConfig.Name,
		PageName: constants.CONVERT_PAGE_NAME,
	}

	return render.Render(c, utils.GetFormatPageTitle(handler.appConfig.Name, options.PageName), pages.Convert(options))
}
