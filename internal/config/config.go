// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.

package config

import (
	"sync"
	"time"
)

const ApplicationName = "girc-ui"

type Config struct {
	mu          sync.RWMutex
	initialSync bool `json:"-"`

	DefaultUser ConfigUser      `json:"default_user"`
	JoinDelay   time.Duration   `json:"join_delay"`
	Servers     []*ConfigServer `json:"servers"`
}

func (c *Config) GetDefaultUser() ConfigUser {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.DefaultUser
}

func (c *Config) SetDefaultUser(user ConfigUser) error {
	c.mu.Lock()
	c.DefaultUser = user
	c.mu.Unlock()

	return c.Write()
}

func (c *Config) GetJoinDelay() time.Duration {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.JoinDelay
}

func (c *Config) SetJoinDelay(delay time.Duration) error {
	c.mu.Lock()
	c.JoinDelay = delay
	c.mu.Unlock()

	return c.Write()
}

func (c *Config) GetServers() []*ConfigServer {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.Servers
}

func (c *Config) GetServer(id string) *ConfigServer {
	c.mu.RLock()
	defer c.mu.RUnlock()

	for _, server := range c.Servers {
		if server.ID == id {
			return server
		}
	}

	return nil
}

func (c *Config) SetServer(server *ConfigServer) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	for i, s := range c.Servers {
		if s.ID == server.ID {
			c.Servers[i] = server
			return c.Write()
		}
	}

	c.Servers = append(c.Servers, server)
	return c.Write()
}

func (c *Config) RemoveServer(id string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	for i, server := range c.Servers {
		if server.ID == id {
			c.Servers = append(c.Servers[:i], c.Servers[i+1:]...)
			return c.Write()
		}
	}

	return nil
}

type ConfigUser struct {
	Nick     string `json:"nick"`
	User     string `json:"user"`
	RealName string `json:"real_name"`
}

type ConfigChannel struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type ConfigServer struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	Host          string `json:"host"`
	Port          int    `json:"port"`
	SSL           bool   `json:"ssl"`
	TLSSkipVerify bool   `json:"tls_skip_verify"`
	DisableSTS    bool   `json:"disable_sts"`
	Password      string `json:"password"`
	LastConnected bool   `json:"last_connected"`

	Channels []*ConfigChannel `json:"channels"`
	User     ConfigUser       `json:"user"`
}
