package models

import "fmt"

type Executor interface {
	init(context Context)
	process(context Context)
	getName() string
}

type Order struct {
	Id        int
	Type      string
	executors []Executor
}

func (o *Order) init(context Context) {
}

func (o *Order) process(context Context) {
	if o.executors != nil {
		for _, executor := range o.executors {
			fmt.Printf("%s: \n", executor.getName())
			executor.process(context)
		}
	}
}

func (o *Order) getName() string {
	return "order"
}

type Uvexecutor struct {
	Id        int
	Type      string
	KeyPrefix string
}

func (o *Uvexecutor) init(context Context) {
}

func (o *Uvexecutor) process(context Context) {
	fmt.Printf("%s: %s \n", o.Type, o.Id)
}

func (o *Uvexecutor) getName() string {
	return "uvexecutor"
}
