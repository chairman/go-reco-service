package models

import "log"

var Ma string

type FetchUpdateRule struct {
	TbName string
}

func (h FetchUpdateRule) Run() {
	Ma = h.TbName
	log.Println(Ma)
}
