package handlers

import (
	"imgix/internal/common/constants"
	"imgix/internal/common/utils"
	"imgix/internal/components/pages"
	"imgix/internal/config"
	"imgix/internal/render"

	"github.com/gofiber/fiber/v2"
)

type ResizeHandler struct {
	appConfig config.AppConfig
}

func NewResizeHandler(appConfig config.AppConfig) *ResizeHandler {
	handler := ResizeHandler{
		appConfig: appConfig,
	}
	return &handler
}

func (handler *ResizeHandler) Index(c *fiber.Ctx) error {
	options := pages.ResizeOptions{
		AppName:  handler.appConfig.Name,
		PageName: constants.RESIZE_PAGE_NAME,
	}

	return render.Render(c, utils.GetFormatPageTitle(handler.appConfig.Name, options.PageName), pages.Resize(options))
}
