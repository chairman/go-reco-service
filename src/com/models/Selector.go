package models

type Selector interface {
	judge(context Context) bool
}
