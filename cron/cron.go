package cron

import (
	"github.com/robfig/cron/v3"
	"log"
)

func NewCron() *cron.Cron {
	c := cron.New()

	id, err := c.AddFunc("* */5 * * * ", func() {
		log.Println("hello world")
	})

	log.Printf("[CRON] id=%d, err=%e", id, err)

	return c
}
