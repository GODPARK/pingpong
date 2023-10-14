package main

import (
	"fmt"
	"log"
	"os"
	"pingpong/client"
	"pingpong/config"
	"pingpong/cron"
	"pingpong/router"
	"time"
)

func main() {

	location, err := time.LoadLocation("Asia/Seoul")
	if err != nil {
		panic(err)
	}
	time.Local = location

	c := config.NewConfig("./config.json")
	client.NewClient(c)

	lf := getLogFile(c.Pong.LogDir)
	defer lf.Close()
	log.SetOutput(lf)

	r := router.NewRouter(c)
	cr := cron.NewCron(c)

	cr.Start()

	log.Printf("[SERVER] start server port=%s", c.Pong.Port)
	r.Run(c.Pong.Port)
}

func getLogFile(path string) *os.File {

	logDirPath := path
	if path == "" {
		logDirPath = "./"
	}

	startDate := time.Now().Format("060102")
	logFilePath := fmt.Sprintf("%s/pingpong_%s.log", logDirPath, startDate)
	if _, err := os.Stat(logFilePath); os.IsNotExist(err) {
		os.MkdirAll(logDirPath, 0777)
	}

	if _, err := os.Stat(logFilePath); os.IsNotExist(err) {
		os.Create(logFilePath)
	}

	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	return logFile
}
