package tasks

import (
	"github.com/robfig/cron"
	"go-reco-service/src/com/models"
	"log"
	"strings"
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
	appName := "wifi"
	resType := "video"
	tbName := getTableName(appName, resType)
	h2 := models.FetchUpdateRule{TbName: tbName}
	c.Schedule(s, h2)
	// 启动任务
	c.Start()
}

func getTableName(appName string, resType string) string {
	var build strings.Builder
	build.WriteString("tb_base_rule_")
	build.WriteString(appName)
	build.WriteString("_")
	build.WriteString(resType)
	tbName := build.String()
	return tbName
}
