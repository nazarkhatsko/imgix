package main

import (
	"imgix/internal/app"
	"imgix/internal/config"
	"imgix/internal/handlers"
	"imgix/internal/routes"

	"go.uber.org/fx"
)

//	@title			IMGIX Swagger API
//	@version		1.0
//	@description	This is a sample server Petstore server.

//	@license.name	MIT
//	@license.url	http://github.com/nazarkhatsko/imgix/LICENSE

//	@BasePath	/api/v1

func main() {
	config.MustLoadConfig()

	fx.New(
		fx.Provide(
			config.MustLoadAppConfig,
			app.InitApp,
		),
		fx.Provide(
			handlers.NewWelcomeHandler,
			handlers.NewConvertHandler,
			handlers.NewCompressHandler,
			handlers.NewResizeHandler,
		),
		fx.Invoke(
			routes.SetupRouteWelcome,
			routes.SetupRouteConvert,
			routes.SetupRouteCompress,
			routes.SetupRouteResize,
			routes.SetupRouteNotFound,
			app.RunApp,
		),
	).Run()
}
