package middlewares

import (
	"imgix/internal/common/constants"
	"imgix/internal/components/pages"
	"imgix/internal/render"

	"github.com/gofiber/fiber/v2"
)

func NotFoundMiddleware(c *fiber.Ctx) error {
	return render.Render(c, constants.NOT_FOUND_PAGE_NAME, pages.NotFound(pages.NotFoundOptions{}))
}
