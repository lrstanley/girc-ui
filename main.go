// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.

package main

import (
	"context"
	"embed"

	"github.com/lrstanley/girc-ui/internal/config"
	"github.com/lrstanley/girc-ui/internal/servermanager"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

var (
	//go:embed all:frontend/dist
	assets embed.FS

	app           *App
	serverManager *servermanager.ServerManager
)

func main() {
	cfg := &config.Config{}

	err := cfg.Sync()
	if err != nil {
		panic(err)
	}

	app = NewApp(cfg)
	serverManager = servermanager.NewServerManager(cfg)

	err = wails.Run(&options.App{
		Title:  config.ApplicationName,
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		Menu:      nil,
		Logger:    nil,
		LogLevel:  logger.DEBUG,
		Frameless: false,
		// Frameless: true,
		OnStartup: func(ctx context.Context) {
			app.Run(ctx)
			serverManager.Run(ctx)
		},
		OnDomReady: app.domReady,
		OnBeforeClose: func(ctx context.Context) bool {
			if app.beforeClose(ctx) {
				return true
			}

			serverManager.Close()
			return false
		},
		OnShutdown:                       app.shutdown,
		WindowStartState:                 options.Normal,
		EnableFraudulentWebsiteDetection: false,
		Bind:                             []any{app, serverManager},
		BackgroundColour:                 &options.RGBA{R: 0, G: 0, B: 0, A: 1},
		// BackgroundColour: &options.RGBA{R: 0, G: 0, B: 0, A: 0},
		// Windows: &windows.Options{
		// 	WindowIsTranslucent: true,
		// },
		// Linux: &linux.Options{
		// 	WindowIsTranslucent: true,
		// },
	})
	if err != nil {
		println("Error:", err.Error())
	}
}
