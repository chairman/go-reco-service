package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-reco-service/src/com/models"
	"io"
	"log"
	"net/http"
	"strings"
)

type Rule struct {
	RuleId, Description string
}

func AddCallbackHandler(c *gin.Context) {
	appName := c.Param("app_name")
	resType := c.Param("res_type")
	rawRes, err := c.GetRawData()
	if err != nil {
		log.Fatal("callback http body is empty")
		return
	}
	if len(rawRes) == 0 {
		log.Fatal("callback http body is empty")
		return
	}
	jsonstr := string(rawRes)
	dec := json.NewDecoder(strings.NewReader(jsonstr))
	var rule Rule
	if err := dec.Decode(&rule); err == io.EOF {
		fmt.Println(err)
	} else if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s: %s \n ", rule.RuleId, rule.Description)
	InsertOneResult := models.NewMgo().InsertOne(rule)
	fmt.Println("Inserted a single document: ", InsertOneResult)
	c.JSON(http.StatusOK, gin.H{"appName": appName, "resType": resType})
}
