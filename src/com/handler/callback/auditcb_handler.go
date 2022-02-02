package callback

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"strings"
)

type Data struct {
	ruleId      int    `json:"ruleId"`
	description string `json:"description"`
}

type Message struct {
	Name, Text string
}

func AuditCallbackHandler(c *gin.Context) {
	engineName := c.Param("engine_name")
	mixSign := c.Param("mix_sign")
	resID := c.Param("res_id")

	rawRes, err := c.GetRawData()
	if err != nil {
		log.Fatal("callback http body is empty")
		return
	}
	if len(rawRes) == 0 {
		log.Fatal("callback http body is empty")
		return
	}
	//buf := make([]byte, 1024)
	//n, _ := c.Request.Body.Read(rawRes)
	jsonstr := string(rawRes)
	//fmt.Println(jsonstr)
	dec := json.NewDecoder(strings.NewReader(jsonstr))
	//for {
	var m Message
	if err := dec.Decode(&m); err == io.EOF {
		fmt.Println(err)
	} else if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s: %s \n ", m.Name, m.Text)
	//}
	//var jsonBlob = []byte(jsonstr)
	//var data Data
	//if err := json.Unmarshal(jsonBlob, &data); err == nil {
	//	fmt.Println(data.ruleId)
	//} else {
	//	fmt.Println(err)
	//}
	c.String(http.StatusOK, "name:%s,id:%s,mixSign:%s", engineName, resID, mixSign)
}
