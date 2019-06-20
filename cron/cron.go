package cron

import (
	"fmt"
	"github.com/robfig/cron"
	"schedule-management/logger"
	"sync"
)

func Run() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	c := cron.New()
	_ = c.AddFunc("1 * * * * *", test)
	c.Start()
	wg.Wait()
}

func test() {
	logger.Info("log cron job", logger.LogFields{})
	fmt.Println("----- cron job ----")
}