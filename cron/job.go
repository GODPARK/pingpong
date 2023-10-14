package cron

import (
	"fmt"
	"net/http"
	"pingpong/client"
	"pingpong/config"
	"pingpong/constants"
	"sync"
	"time"
)

var Job JobResult

func init() {
	Job.Result = make(map[string]JobResultDetail)
}

type JobResult struct {
	Result map[string]JobResultDetail
	mu     sync.Mutex
}

type JobResultDetail struct {
	Date      []string
	Status    []bool
	Msg       []string
	FailCount int
}

func setJob(name string, detail JobResultDetail) {
	Job.mu.Lock()
	defer Job.mu.Unlock()

	if name == "" {
		return
	}

	Job.Result[name] = detail
}

func getJob(name string) JobResultDetail {
	Job.mu.Lock()
	defer Job.mu.Unlock()

	j, ok := Job.Result[name]
	if !ok {
		return JobResultDetail{}
	}
	return j
}

func pingJob(config *config.Config) {
	for _, v := range config.Ping.Hosts {
		j := getJob(v.Name)
		status, msg := requestPing(v)
		if status != v.Status {
			j.FailCount++
			j.Status = append(j.Status, false)
		} else {
			j.FailCount = 0
			j.Status = append(j.Status, true)
		}

		j.Msg = append(j.Msg, msg)
		j.Date = append(j.Date, time.Now().String())
		setJob(v.Name, j)
	}
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
