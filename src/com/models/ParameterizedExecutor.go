package models

type ParameterizedExecutor interface {
	megre(context Context) Executor
}
