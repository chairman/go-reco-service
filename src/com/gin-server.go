package main

import (
	_ "go-reco-service/src/com/drivers"
	"go-reco-service/src/com/http"
	"go-reco-service/src/com/redis"
	"go-reco-service/src/com/tasks"
)

func main() {
	//drivers.Init()
	redis.Init()
	tasks.Init()
	http.Lanuch("0.0.0.0:8658")
}
