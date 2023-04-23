package main

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/lrstanley/girc"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context

	mu      sync.RWMutex
	clients []*girc.Client
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

type LogWriter struct {
	ctx context.Context
}

func (l *LogWriter) Write(p []byte) (n int, err error) {
	runtime.LogDebug(l.ctx, strings.TrimSuffix(string(p), "\n"))
	return len(p), nil
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	// runtime.Hide(ctx)

	a.clients = []*girc.Client{
		girc.New(girc.Config{
			Server:  "localhost",
			Port:    6667,
			Nick:    "girc-ui-1",
			User:    "girc-ui-1",
			Name:    "girc-ui-1",
			Version: "girc-ui debug version",
			Debug:   &LogWriter{ctx},
		}),
		// girc.New(girc.Config{
		// 	Server:  "localhost",
		// 	Port:    6667,
		// 	Nick:    "girc-ui-2",
		// 	User:    "girc-ui-2",
		// 	Name:    "girc-ui-2",
		// 	Version: "girc-ui debug version",
		// 	Debug:   &LogWriter{ctx},
		// }),
		// girc.New(girc.Config{
		// 	Server:  "localhost",
		// 	Port:    6667,
		// 	Nick:    "girc-ui-3",
		// 	User:    "girc-ui-3",
		// 	Name:    "girc-ui-3",
		// 	Version: "girc-ui debug version",
		// 	Debug:   &LogWriter{ctx},
		// }),
	}

	// a.mu.RLock()
	// defer a.mu.RUnlock()

	for _, client := range a.clients {
		go func(c *girc.Client) {
			// Register all preconfgured channels, and also forward events to the UI.
			c.Handlers.AddBg(girc.CONNECTED, func(c *girc.Client, e girc.Event) {
				c.Cmd.Join("#girc-ui")
			})

			c.Handlers.AddBg(girc.ALL_EVENTS, func(c *girc.Client, e girc.Event) {
				runtime.EventsEmit(ctx, "IRC_"+e.Command, e)
				runtime.EventsEmit(ctx, "IRC_ALL", e)
			})

			c.Handlers.AddBg(girc.UPDATE_STATE, func(c *girc.Client, e girc.Event) {
				runtime.EventsEmit(ctx, "irc-update-state", e)
			})

			time.Sleep(5 * time.Second)

			for {
				err := c.Connect()
				if err == nil {
					return
				}

				fmt.Printf("error: %s", err)

				fmt.Println("disconnected, reconnecting in 10 seconds")
				time.Sleep(10 * time.Second)
			}
		}(client)
	}
}

// domReady is called after front-end resources have been loaded
func (a *App) domReady(ctx context.Context) {
	// Add your action here
}

// beforeClose is called when the application is about to quit,
// either by clicking the window close button or calling runtime.Quit.
// Returning true will cause the application to continue, false will continue shutdown as normal.
func (a *App) beforeClose(ctx context.Context) (prevent bool) {
	a.mu.Lock()
	defer a.mu.Unlock()

	for _, client := range a.clients {
		client.Quit("closing client")
	}

	time.Sleep(500 * time.Millisecond)

	return false
}

// shutdown is called at application termination
func (a *App) shutdown(ctx context.Context) {
	// Perform your teardown here
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) GetUsersForChannel(channel string) []*girc.User {
	// a.mu.RLock()
	// defer a.mu.RUnlock()

	for _, client := range a.clients {
		ch := client.LookupChannel(channel)
		if ch == nil {
			continue
		}

		return ch.Users(client)
	}

	return nil
}
