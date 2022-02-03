package tasks

import (
	"github.com/robfig/cron"
	"go-reco-service/src/com/models"
	"log"
)

func Init() {
	Service()
}

func Service() {
	c := cron.New()
	s, err := cron.Parse("*/3 * * * * *")
	if err != nil {
		log.Println("Parse error")
	}
	h2 := models.Hello{Str: "I Hate You!"}
	c.Schedule(s, h2)
	// 启动任务
	c.Start()
}
