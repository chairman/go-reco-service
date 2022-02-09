package models

import "fmt"

type Executor interface {
	Init(context Context)
	Process(context Context)
	GetName() string
}

type Order struct {
	Id        int
	Type      string
	executors []Executor
}

func (o *Order) Init(context Context) {
}

func (o *Order) Process(context Context) {
	if o.executors != nil {
		for _, executor := range o.executors {
			fmt.Printf("%s: \n", executor.GetName())
			executor.Process(context)
		}
	}
}

func (o *Order) GetName() string {
	return "order"
}

type Uvexecutor struct {
	Id        int
	Type      string
	KeyPrefix string
}

func (o *Uvexecutor) Init(context Context) {
}

func (o *Uvexecutor) Process(context Context) {
	fmt.Printf("%s: %d \n", o.Type, o.Id)
}

func (o *Uvexecutor) GetName() string {
	return "uvexecutor"
}
