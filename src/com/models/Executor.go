package models

import (
	"fmt"
	"go-reco-service/src/com/redis"
)

type Executor interface {
	Process(context Context) []int
	GetName() string
}

type Order struct {
	Id        int
	Type      string
	executors []Executor
}

func (o *Order) Process(context Context) []int {
	var result []int
	if o.executors != nil {
		for _, executor := range o.executors {
			fmt.Printf("%s: \n", executor.GetName())
			var t1 = result
			var t2 = executor.Process(context)

			var r []int

			for _, v := range t1 {
				r = append(r, v)
			}

			for _, v := range t2 {
				r = append(r, v)
			}

			result = r
		}
	}
	return result
}

func (o *Order) GetName() string {
	return "order"
}

type Uvexecutor struct {
	Id        int
	Type      string
	KeyPrefix string
}

func (o *Uvexecutor) Process(context Context) []int {
	fmt.Printf("%s: %d \n", o.Type, o.Id)
	//recoData := context.RecoData
	//RecoData = append(recoData, 1)
	var result []int
	result = append(result, 1)
	redis.Set("b", "222")
	return result
}

func (o *Uvexecutor) GetName() string {
	return "uvexecutor"
}
