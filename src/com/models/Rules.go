package models

import (
	"context"
	"fmt"
	"go-reco-service/src/com/utils"
	"log"
	"strconv"
)

type FetchUpdateRule struct {
	TbName string
}

func (fetchUpdateRule FetchUpdateRule) Run() {
	RuleConfigs = getRuleConfigs(fetchUpdateRule)
	var rs []*Rule
	for _, config := range RuleConfigs {
		selectorConfigStr := config.Selector
		exectorConfigStr := config.Executor
		fmt.Printf("%s: %s \n", selectorConfigStr, exectorConfigStr)
		var rule = new(Rule)
		rule.RuleId = config.RuleId
		rule.Order = config.Order
		rule.Description = config.Description
		rule.OP = config.OP
		selectorNode, err := utils.LoadByString(selectorConfigStr)
		if err == nil {
			rule.Selector = ParseSelector(selectorNode)
		}
		exectorNode, err := utils.LoadByString(exectorConfigStr)
		if err == nil {
			rule.Executor = ParseExecutor(exectorNode)
			if rule.Executor == nil {
				log.Println("update rule err:", config.RuleId)
			} else {
				rs = append(rs, rule)
			}
		}
	}
	Rules = rs
}

func ParseSelector(selectorNode *utils.JsonNode) Selector {
	dats := selectorNode.StructNodes
	for selectorType, _ := range dats {
		fmt.Println("name:", selectorType)
		selector := CreateSelector(selectorType, selectorNode)
		if selector != nil {
			fmt.Println(" selector.getName:", selector.getName())
			return selector
		}
	}
	return nil
}

func CreateSelector(selectorType string, selectorNode *utils.JsonNode) Selector {
	switch selectorType {
	case "$true":
		return &True{}
	case "$eq":
		nodes := selectorNode.StructNodes[selectorType]
		arraysStruct := nodes.ToJsonNode().ArraysStruct
		nodesSize := len(arraysStruct)
		fmt.Println(" selector.nodesSize:", nodesSize)
		if nodesSize == 2 {
			key := arraysStruct[0].ToJsonNode().ValueString
			value := int(arraysStruct[1].ToJsonNode().ValueNumber)
			fmt.Println(" selector.key:", key)
			fmt.Println(" selector.value:", value)
			return &Eq{key, value}
		}
		return nil
	case "$and":
		nodes := selectorNode.StructNodes[selectorType]
		arraysStruct := nodes.ToJsonNode().ArraysStruct
		nodesSize := len(arraysStruct)
		fmt.Println(" selector.nodesSize:", nodesSize)
		fmt.Println(" executor.getName:", nodesSize)
		var selectors = make([]Selector, nodesSize)
		for i := 0; i < len(arraysStruct); i++ {
			selectors[i] = ParseSelector(arraysStruct[i].ToJsonNode())
		}
		return &And{selectors}
	case "$or":
		nodes := selectorNode.StructNodes[selectorType]
		arraysStruct := nodes.ToJsonNode().ArraysStruct
		nodesSize := len(arraysStruct)
		fmt.Println(" selector.nodesSize:", nodesSize)
		fmt.Println(" executor.getName:", nodesSize)
		var selectors = make([]Selector, nodesSize)
		for i := 0; i < len(arraysStruct); i++ {
			selectors[i] = ParseSelector(arraysStruct[i].ToJsonNode())
		}
		return &Or{selectors}
	case "$gt":
		return &Gt{}
	case "$gte":
		return &Gte{}
	case "$lt":
		return &Lt{}
	case "$lte":
		return &Lte{}
	default:
		return nil
	}
}

func ParseExecutor(exectorNode *utils.JsonNode) Executor {
	excutorType := exectorNode.GetNodeByPath("type").ValueString
	excutorId, err := strconv.Atoi(exectorNode.GetNodeByPath("id").ValueString)
	fmt.Println("excutorType:", excutorType)
	fmt.Println("excutorId:", excutorId)
	if excutorType == "" || err != nil {
		return nil
	}
	executor := CreateExecutor(excutorId, excutorType, exectorNode)
	fmt.Println(" executor.getName:", executor.GetName())
	return executor
}

func CreateExecutor(excutorId int, executorType string, node *utils.JsonNode) Executor {
	switch executorType {
	case "order":
		children := node.GetNodeByPath("children").ArraysStruct
		childrenLen := len(children)
		fmt.Println(" executor.getName:", childrenLen)
		var executors = make([]Executor, childrenLen)
		for i := 0; i < len(children); i++ {
			executors[i] = ParseExecutor(children[i].ToJsonNode())
		}
		return &Order{excutorId, executorType, executors}
	case "uvexecutor":
		return &Uvexecutor{excutorId, executorType, node.GetNodeByPath("keyPrefix").ValueString}
	default:
		return nil
	}
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
