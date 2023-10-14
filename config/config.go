package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	Ping struct {
		Cron  string `json:"cron"`
		Hosts []Host `json:"hosts"`
	} `json:"ping"`
	Pong struct {
		Port   string `json:"port"`
		Status int    `json:"status"`
		LogDir string `json:"logDir"`
		Token  string `json:"token"`
	} `json:"pong"`
	Alert struct {
		Slack struct {
		} `json:"slack"`
		Email struct {
		} `json:"email"`
	} `json:"alert"`
}

type Host struct {
	Name           string `json:"name"`
	Target         string `json:"host"`
	Token          string `json:"token"`
	Status         int    `json:"status"`
	Timeout        int    `json:"timeout"`
	ErrorThreshold int    `json:"error_threshold"`
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

	if c.Ping.Cron == "" {
		panic("cron expression is empty")
	}

	return c
}
