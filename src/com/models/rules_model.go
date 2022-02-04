package models

import (
	"context"
	"fmt"
	"log"
)

type FetchUpdateRule struct {
	TbName string
}

func (fetchUpdateRule FetchUpdateRule) Run() {
	RuleConfigs = getRuleConfigs(fetchUpdateRule)
	var rs []*Rule
	for _, config := range RuleConfigs {
		selectorConfig := config.Selector
		exectorConfig := config.Executor
		fmt.Printf("%s: %s \n", selectorConfig, exectorConfig)
		var rule = new(Rule)
		rule.RuleId = config.RuleId
		rule.Order = config.Order
		rule.Description = config.Description
		rule.OP = config.OP
		rs = append(rs, rule)
	}
	Rules = rs
}

var RuleConfigs []*RuleConfig

var Rules []*Rule

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
