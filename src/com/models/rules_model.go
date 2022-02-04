package models

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
)

type FetchUpdateRule struct {
	TbName string
}

func (h FetchUpdateRule) Run() {
	RuleConfigs = getRuleConfigs(h)
	jsons, errs := json.Marshal(RuleConfigs) //转换成JSON返回的是byte[]
	if errs != nil {
		fmt.Println(errs.Error())
	}
	fmt.Printf("Found multiple documents (array of pointers): %+v\n", string(jsons))

}

var RuleConfigs []*RuleConfig

type Rule struct {
	RuleId      int
	Order       int
	Description string
	OP          string
	Selector    Selector
	executor    ParameterizedExecutor
}

func getRuleConfigs(h FetchUpdateRule) []*RuleConfig {
	cur := NewMgo(h.TbName).FindAll(0, 1000, 1)
	var rs []*RuleConfig
	for cur.Next(context.TODO()) {
		// create a value into which the single document can be decoded
		var elem RuleConfig
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		rs = append(rs, &elem)
	}
	RuleConfigs = rs

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// Close the cursor once finished
	err := cur.Close(context.TODO())
	if err != nil {
		return nil
	}
	return RuleConfigs
}
