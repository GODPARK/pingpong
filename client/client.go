package client

import (
	"net/http"
	"pingpong/config"
	"sync"
	"time"
)

var GlobalClient Client

type Client struct {
	TargetClient map[string]*http.Client
	mu           sync.Mutex
}

func NewClient(config *config.Config) {
	GlobalClient.TargetClient = make(map[string]*http.Client)

	for _, v := range config.Ping.Hosts {
		if v.Name != "" && v.Target != "" {
			h := &http.Client{}
			if v.Timeout == 0 {
				h.Timeout = time.Duration(1000) * time.Millisecond
			} else {
				h.Timeout = time.Duration(v.Timeout) * time.Millisecond
			}
			SetGlobalClient(v.Name, h)
		}
	}
}

func SetGlobalClient(name string, h *http.Client) {
	GlobalClient.mu.Lock()
	defer GlobalClient.mu.Unlock()

	GlobalClient.TargetClient[name] = h
}

func GetGlobalClient(name string) (*http.Client, bool) {
	GlobalClient.mu.Lock()
	defer GlobalClient.mu.Unlock()

	c, ok := GlobalClient.TargetClient[name]

	return c, ok
}
