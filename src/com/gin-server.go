package main

import (
	"go-reco-service/src/com/drivers"
	"go-reco-service/src/com/http"
)

func main() {
	drivers.Init()
	http.Lanuch("0.0.0.0:8658")
}
