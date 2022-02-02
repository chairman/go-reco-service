package callback

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

type Message struct {
	RuleId, Description string
}

//type Message struct {
//	Name, Text string
//}
type Trainer struct {
	Name string
	Age  int
	City string
}

func AuditCallbackHandler(c *gin.Context) {
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
	var m Message
	if err := dec.Decode(&m); err == io.EOF {
		fmt.Println(err)
	} else if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s: %s \n ", m.RuleId, m.Description)
	// 单个插入
	//ash := Trainer{"Ash", 10, "Pallet Town"}
	InsertOneResult := models.NewMgo().InsertOne(m)
	fmt.Println("Inserted a single document: ", InsertOneResult)
	c.JSON(http.StatusOK, gin.H{"appName": appName, "resType": resType})
}
