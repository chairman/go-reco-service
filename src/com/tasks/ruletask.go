package tasks

import (
	"github.com/robfig/cron"
	"go-reco-service/src/com/models"
	"log"
)

//type Ma struct {
//	Str string
//}
//
//var ma Ma

//type Hello struct {
//	Str string
//}
//
//func (h Hello) Run() {
//	log.Println(h.Str)
//	ma := h.Str
//	log.Println(ma)
//}

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
	// 关闭任务
	//defer c.Stop()
	//select {}
}
