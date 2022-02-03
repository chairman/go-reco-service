package models

import "log"

var Ma string

type Hello struct {
	Str string
}

func (h Hello) Run() {
	Ma = h.Str
	log.Println(Ma)
}
