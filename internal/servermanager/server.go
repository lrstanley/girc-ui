// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.

package servermanager

import (
	"context"
	"crypto/tls"
	"sync"
	"time"

	"github.com/lrstanley/girc"
	"github.com/lrstanley/girc-ui/internal/config"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type Server struct {
	id     string
	ctx    context.Context
	cfg    *config.Config
	client *girc.Client

	mu     sync.RWMutex
	events []girc.Event
}

func NewServer(ctx context.Context, cfg *config.Config, id string) *Server {
	serverCfg := cfg.GetServer(id)

	ircConfig := girc.Config{
		Server:     serverCfg.Host,
		ServerPass: serverCfg.Password,
		Port:       serverCfg.Port,
		SSL:        serverCfg.SSL,
		TLSConfig:  &tls.Config{InsecureSkipVerify: serverCfg.TLSSkipVerify},
		DisableSTS: serverCfg.DisableSTS,
		Debug:      &LogWriter{ctx},
		Nick:       serverCfg.User.Nick,
		User:       serverCfg.User.User,
		Name:       serverCfg.User.RealName,
	}

	defaultUser := cfg.GetDefaultUser()

	if ircConfig.Nick == "" {
		ircConfig.Nick = defaultUser.Nick
	}

	if ircConfig.User == "" {
		ircConfig.User = defaultUser.User
	}

	if ircConfig.Name == "" {
		ircConfig.Name = defaultUser.RealName
	}

	client := girc.New(ircConfig)
	server := &Server{id: id, ctx: ctx, cfg: cfg, client: client}

	client.Handlers.AddBg(girc.ALL_EVENTS, server.handleEvent)

	go server.Connect()
	return server
}

func (s *Server) handleEvent(c *girc.Client, e girc.Event) {
	switch e.Command {
	case girc.CONNECTED:
		serverCfg := s.cfg.GetServer(s.id)
		joinDelay := s.cfg.GetJoinDelay()

		for _, channel := range serverCfg.Channels {
			if channel.Password != "" {
				c.Cmd.JoinKey(channel.Name, channel.Password)
			} else {
				c.Cmd.Join(channel.Name)
			}

			if joinDelay > 0 {
				time.Sleep(joinDelay)
			}
		}
	case girc.UPDATE_STATE:
		runtime.EventsEmit(s.ctx, "irc-update-state", e)
	case girc.UPDATE_GENERAL:
		runtime.EventsEmit(s.ctx, "irc-update-general", e)
	default:
		s.mu.Lock()
		s.events = append(s.events, e)
		s.mu.Unlock()

		runtime.EventsEmit(s.ctx, "IRC_"+e.Command, e)
		runtime.EventsEmit(s.ctx, "IRC_ALL", e)
	}
}

func (s *Server) Connect() {
	for {
		err := s.client.Connect()
		if err == nil {
			return
		}

		runtime.LogErrorf(s.ctx, "error: %s", err)
		runtime.LogError(s.ctx, "disconnected, reconnecting in 10 seconds")
		time.Sleep(10 * time.Second)
	}
}

func (s *Server) Close() {
	s.client.Quit("closing client")
}
