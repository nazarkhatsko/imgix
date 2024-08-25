package render

import (
	"imgix/internal/components/layouts"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

func Render(c *fiber.Ctx, title string, embed templ.Component) error {
	handler := templ.Handler(layouts.Main(layouts.MainOptions{
		Lang:  "en",
		Title: title,
		Embed: embed,
	}))

	return adaptor.HTTPHandler(handler)(c)
}
