package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-reco-service/src/com/models"
	"go.mongodb.org/mongo-driver/bson"
	"io"
	"log"
	"net/http"
	"strings"
)

type Rule struct {
	RuleId, Order, Description, OP, Selector, Executor string
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

func AddHandler(c *gin.Context) {
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
	tbName := getTableName(appName, resType)
	InsertOneResult := models.NewMgo(tbName).InsertOne(rule)
	fmt.Println("Inserted a single document: ", InsertOneResult)
	c.JSON(http.StatusOK, gin.H{"msg": "ok", "code": 200})
}

func UpdateHandler(c *gin.Context) {
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
	tbName := getTableName(appName, resType)
	data := bson.M{"$set": bson.M{"order": rule.Order, "description": rule.Description, "op": rule.OP, "selector": rule.Selector, "executor": rule.Executor}}
	updateResult := models.NewMgo(tbName).UpdateOne("ruleid", rule.RuleId, data)
	fmt.Println("Inserted a single document: ", updateResult)
	c.JSON(http.StatusOK, gin.H{"appName": appName, "resType": resType})
}

func DeleteHandler(c *gin.Context) {
	appName := c.Param("app_name")
	resType := c.Param("res_type")
	ruleId := c.Param("rule_id")
	tbName := getTableName(appName, resType)
	updateResult := models.NewMgo(tbName).Delete("ruleid", ruleId)
	fmt.Println("Inserted a single document: ", updateResult)
	c.JSON(http.StatusOK, gin.H{"appName": appName, "resType": resType})
}
