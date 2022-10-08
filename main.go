package main

import (
	"github.com/robfig/cron/v3"
	"time"
	"fmt"
	"twitter-bot/internal"
)

func main() {
	c := cron.New()
	c.AddFunc("@every 1m", internal.Job)
	c.Start()
	for {
		time.Sleep(10000000000000)
		fmt.Println("sleep")
	}
}