package models

type RuleConfig struct {
	RuleId      int
	Order       int
	Description string
	OP          string
	Selector    string
	Executor    string
}
