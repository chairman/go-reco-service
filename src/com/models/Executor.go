package models

type Executor interface {
	init(context Context)
	process(context Context)
	getId() string
	getType() string
	cleanup(context Context)
}
