package tasks

import (
	"github.com/robfig/cron"
	"log"
)

type Hello struct {
	Str string
}

func (h Hello) Run() {
	log.Println(h.Str)
}

func Init() {
	Service()
}

func Service() {
	c := cron.New()
	s, err := cron.Parse("*/3 * * * * *")
	if err != nil {
		log.Println("Parse error")
	}
	h2 := Hello{"I Hate You!"}
	c.Schedule(s, h2)
	// 启动任务
	c.Start()
	// 关闭任务
	//defer c.Stop()
	//select {}
}
