package cron

import (
	"github.com/robfig/cron/v3"
	"log"
	"pingpong/config"
)

func NewCron(config *config.Config) *cron.Cron {
	c := cron.New()

	id, err := c.AddFunc(config.Ping.Cron, func() {
		pingJob(config)
	})
	if err != nil {
		panic(err)
	}

	log.Printf("[CRON] id=%d", id)

	return c
}
