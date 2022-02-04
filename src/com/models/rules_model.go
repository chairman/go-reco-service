package models

import (
	"context"
	"fmt"
	"log"
)

var Ma string

type FetchUpdateRule struct {
	TbName string
}

type Rule struct {
	RuleId, Order, Description, OP, Selector, Executor string
}

var Results []*Rule

func (h FetchUpdateRule) Run() {
	Ma = h.TbName
	cur := NewMgo(h.TbName).FindAll(0, 1000, 1)
	var rs []*Rule
	for cur.Next(context.TODO()) {
		// create a value into which the single document can be decoded
		var elem Rule
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		rs = append(rs, &elem)
	}
	Results = rs

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// Close the cursor once finished
	cur.Close(context.TODO())

	fmt.Printf("Found multiple documents (array of pointers): %+v\n", Results)
	log.Println(Ma)
}
