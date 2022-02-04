package models

import (
	"context"
	"fmt"
	"log"
)

type FetchUpdateRule struct {
	TbName string
}

var RuleConfigs []*RuleConfig

type Rule struct {
	RuleId, Order, Description, OP string
	Selector                       Selector
	executor                       ParameterizedExecutor
}

func (h FetchUpdateRule) Run() {
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
		return
	}

	fmt.Printf("Found multiple documents (array of pointers): %+v\n", RuleConfigs)
}
