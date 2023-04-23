package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed all:frontend/dist
var assets embed.FS

var applicationName = "girc-ui"

func main() {
	app := NewApp()

	err := wails.Run(&options.App{
		Title:  applicationName,
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		Menu:                             nil,
		Logger:                           nil,
		LogLevel:                         logger.DEBUG,
		Frameless:                        true,
		OnStartup:                        app.startup,
		OnDomReady:                       app.domReady,
		OnBeforeClose:                    app.beforeClose,
		OnShutdown:                       app.shutdown,
		WindowStartState:                 options.Normal,
		EnableFraudulentWebsiteDetection: false,
		Bind:                             []any{app},
		// BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		BackgroundColour: &options.RGBA{R: 0, G: 0, B: 0, A: 0},
		Windows: &windows.Options{
			WindowIsTranslucent: true,
		},
		Linux: &linux.Options{
			WindowIsTranslucent: true,
		},
	})
	if err != nil {
		println("Error:", err.Error())
	}
}
