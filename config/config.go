package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	Ping struct {
		Hosts []struct {
			Name   string `json:"name"`
			Host   string `json:"host"`
			Port   string `json:"port"`
			Path   string `json:"path"`
			Status int    `json:"status"`
		} `json:"hosts"`
	} `json:"ping"`
	Pong struct {
		Port   string `json:"port"`
		Status int    `json:"status"`
		LogDir string `json:"logDir"`
	} `json:"pong"`
	Alert struct {
		Slack struct {
		} `json:"slack"`
		Email struct {
		} `json:"email"`
	} `json:"alert"`
}

func NewConfig(path string) *Config {
	c := &Config{}

	log.Printf("[SERVER] config path = %s", path)

	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal(data, c); err != nil {
		panic(err)
	}

	if c.Pong.Port == "" {
		c.Pong.Port = ":9991"
	}

	if c.Pong.Status == 0 {
		c.Pong.Status = 200
	}

	return c
}
