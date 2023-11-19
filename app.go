// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.

package main

import (
	"context"
	"fmt"

	"github.com/lrstanley/girc-ui/internal/config"
)

// App struct
type App struct {
	ctx context.Context
	cfg *config.Config
}

// NewApp creates a new App application struct
func NewApp(cfg *config.Config) *App {
	return &App{cfg: cfg}
}

// Run is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) Run(ctx context.Context) {
	a.ctx = ctx
}

// domReady is called after front-end resources have been loaded.
func (a *App) domReady(ctx context.Context) {}

// beforeClose is called when the application is about to quit,
// either by clicking the window close button or calling runtime.Quit.
// Returning true will cause the application to continue, false will continue shutdown as normal.
func (a *App) beforeClose(ctx context.Context) (prevent bool) {
	return false
}

// shutdown is called at application termination.
func (a *App) shutdown(ctx context.Context) {
}

// Greet returns a greeting for the given name.
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}
