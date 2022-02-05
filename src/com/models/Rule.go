package models

type Rule struct {
	RuleId      int
	Order       int
	Description string
	OP          string
	Selector    Selector
	Executor    ParameterizedExecutor
}
