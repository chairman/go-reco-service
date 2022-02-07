package models

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

func (Order) init(context Context) {
}

func (Order) process(context Context) {
}

func (Order) getName() string {
	return "order"
}

type Uvexecutor struct {
	Id        int
	Type      string
	KeyPrefix string
}

func (Uvexecutor) init(context Context) {
}

func (Uvexecutor) process(context Context) {
}

func (Uvexecutor) getName() string {
	return "uvexecutor"
}
