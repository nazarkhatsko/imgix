package handlers

import (
	"imgix/internal/common/constants"
	"imgix/internal/common/utils"
	"imgix/internal/components/pages"
	"imgix/internal/config"
	"imgix/internal/render"

	"github.com/gofiber/fiber/v2"
)

type CompressHandler struct {
	appConfig config.AppConfig
}

func NewCompressHandler(appConfig config.AppConfig) *CompressHandler {
	handler := CompressHandler{
		appConfig: appConfig,
	}
	return &handler
}

func (handler *CompressHandler) Index(c *fiber.Ctx) error {
	options := pages.CompressOptions{
		AppName:  handler.appConfig.Name,
		PageName: constants.COMPRESS_PAGE_NAME,
	}

	return render.Render(c, utils.GetFormatPageTitle(handler.appConfig.Name, options.PageName), pages.Compress(options))
}
