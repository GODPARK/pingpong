package cron

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"log"
	"net/http"
	"pingpong/client"
	"pingpong/config"
	"pingpong/constants"
)

type JobResult struct {
	Status int
}

type JobStatus struct {
}

func NewCron(config *config.Config) *cron.Cron {
	c := cron.New()

	id, err := c.AddFunc(config.Ping.Cron, func() {
		//for _, v := range config.Ping.Hosts {
		//
		//}
	})
	if err != nil {
		panic(err)
	}

	log.Printf("[CRON] id=%d", id)

	return c
}

func requestPing(host config.Host) (int, string) {
	if host.Name == "" || host.Target == "" {
		return -1, "target, name is empty"
	}

	req, err1 := http.NewRequest("GET", host.Target, nil)
	if err1 != nil {
		return -1, fmt.Sprintf("[%s][%s] request build error=%e", host.Name, host.Target, err1)
	}

	if host.Token != "" {
		req.Header.Add(constants.PingPongTokenHeaderKey, host.Token)
	}
	req.Header.Add("User-Agent", "PingPong")

	c, ok := client.GetGlobalClient(host.Name)
	if !ok {
		return -1, fmt.Sprintf("http client is not set")
	}

	resp, err2 := c.Do(req)
	if err2 != nil {
		return -1, fmt.Sprintf("[%s][%s] request do is error=%e", host.Name, host.Target, err2)
	}
	defer resp.Body.Close()

	return resp.StatusCode, ""

}
