package main

import (
	"sync/atomic"
	"time"
)

type Config struct {
	FeatureFlagA bool
	FeatureFlagB bool
	FeatureFlagC bool
}

type ConfigManager struct {
	data atomic.Value
	stopCh chan struct{}
}

func NewConfigManager() *ConfigManager {
	out := &ConfigManager{
		stopCh: make(chan struct{}),
	}

	out.data.Store(&Config{})

	go out.monitorConfig()

	return out
}

func (c *ConfigManager) monitorConfig() {
	ticker := time.NewTicker(30 * time.Second)

	for {
		select {
		case <- ticker.C:
			cfg := loadConfigFromDatastore()
			c.data.Store(cfg)

		case <- c.stopCh:
			return
		}
	}
}

func (c *ConfigManager) Config() *Config {
	return c.data.Load().(*Config)
}

func (c *ConfigManager) ShutDown() {
	close(c.stopCh)
}

func loadConfigFromDatastore() *Config {
	return &Config{}
}