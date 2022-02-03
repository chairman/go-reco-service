package main

import (
	"log"

	"github.com/robfig/cron"
)

func main() {
	log.Println("Starting...")
	c := cron.New() // 新建一个定时任务对象
	c.AddFunc("* * * * * *", func() {
		log.Println("hello world")
	}) // 给对象增加定时任务
	c.Start()
	select {}
}
