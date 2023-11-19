// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.

package config

import (
	"encoding/json"
	"os"
	"path"
)

func (c *Config) GetConfigFile() (string, error) {
	dir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	dir = path.Join(dir, ApplicationName)
	err = os.MkdirAll(dir, 0o750)
	if err != nil {
		return "", err
	}

	return path.Join(dir, "config.json"), nil
}

func (c *Config) Read() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	configPath, err := c.GetConfigFile()
	if err != nil {
		return err
	}

	f, err := os.Open(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}

		return err
	}
	defer f.Close()

	return json.NewDecoder(f).Decode(c)
}

func (c *Config) Write() error {
	c.mu.RLock()
	defer c.mu.RUnlock()

	configPath, err := c.GetConfigFile()
	if err != nil {
		return err
	}

	f, err := os.Create(configPath)
	if err != nil {
		return err
	}
	defer f.Close()

	return json.NewEncoder(f).Encode(c)
}

func (c *Config) Sync() error {
	c.mu.Lock()

	if c.initialSync {
		c.mu.Unlock()
		return c.Write()
	}

	c.initialSync = true
	c.mu.Unlock()
	if err := c.Read(); err != nil {
		return err
	}

	return c.Write()
}
