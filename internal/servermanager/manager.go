// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.

package servermanager

import (
	"context"
	"sync"
	"time"

	"github.com/lrstanley/girc"
	"github.com/lrstanley/girc-ui/internal/config"
)

type ServerManager struct {
	ctx context.Context
	cfg *config.Config

	mu      sync.RWMutex
	servers map[string]*Server
}

func NewServerManager(cfg *config.Config) *ServerManager {
	return &ServerManager{cfg: cfg, servers: make(map[string]*Server)}
}

func (sm *ServerManager) Run(ctx context.Context) {
	sm.ctx = ctx

	sm.mu.Lock()
	defer sm.mu.Unlock()

	for _, server := range sm.cfg.GetServers() {
		sm.servers[server.ID] = NewServer(sm.ctx, sm.cfg, server.ID)
	}
}

func (sm *ServerManager) Close() {
	sm.mu.RLock()
	defer sm.mu.RUnlock()

	for _, server := range sm.servers {
		server.Close()
	}

	time.Sleep(500 * time.Millisecond)
}

func (sm *ServerManager) GetServersIDs() []string {
	sm.mu.RLock()
	defer sm.mu.RUnlock()

	var ids []string
	for id := range sm.servers {
		ids = append(ids, id)
	}

	return ids
}

func (sm *ServerManager) GetServerConfig(serverID string) *config.ConfigServer {
	return sm.cfg.GetServer(serverID)
}

func (sm *ServerManager) GetServerEvents(serverID string) []girc.Event {
	sm.mu.RLock()
	defer sm.mu.RUnlock()

	s, ok := sm.servers[serverID]
	if !ok {
		return nil
	}

	return s.events
}

func (sm *ServerManager) GetServerChannels(serverID string) []*girc.Channel {
	sm.mu.RLock()
	defer sm.mu.RUnlock()

	s, ok := sm.servers[serverID]
	if !ok {
		return nil
	}

	return s.client.Channels()
}

func (sm *ServerManager) GetServerChannel(serverID, channel string) *girc.Channel {
	sm.mu.RLock()
	defer sm.mu.RUnlock()

	s, ok := sm.servers[serverID]
	if !ok {
		return nil
	}

	return s.client.LookupChannel(channel)
}

func (sm *ServerManager) GetUsersForChannel(serverID, channel string) []*girc.User {
	sm.mu.RLock()
	defer sm.mu.RUnlock()

	s, ok := sm.servers[serverID]
	if !ok {
		return nil
	}

	ch := s.client.LookupChannel(channel)
	if ch == nil {
		return nil
	}

	return ch.Users(s.client)
}
