package models

import (
	"fmt"
	"go-reco-service/src/com/utils"
	"sync"
)

type Registry struct {
	ExecutorMap map[string]ParameterizedExecutor
	SelectorMap map[string]Selector
}

//建立私有变量
var instance *Registry

var once sync.Once

func GetRegistryInstance() *Registry {
	once.Do(func() {
		instance = new(Registry)
	})
	return instance
}

func (registry *Registry) ParseSelector(selectorConfig string) Selector {
	fmt.Println("selectorConfig jsonstr:", selectorConfig)
	return nil
}

func (registry *Registry) ParseExecutor(exectorConfigStr string) ParameterizedExecutor {
	fmt.Println("exectorConfig jsonstr:", exectorConfigStr)
	node, err := utils.LoadByString(exectorConfigStr)

	if err == nil {
		excutorType := node.GetNodeByPath("type")
		fmt.Println("excutorType:", excutorType.ValueString)
		if excutorType.ValueString == "" {
			return nil
		}
		//return GetRegistryInstance().ExecutorMap["type"]
	}
	return nil
}
