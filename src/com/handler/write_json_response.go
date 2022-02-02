package handler

import (
	"encoding/json"
	"fmt"
	"go-reco-service/src/com/models"
	"net/http"
)

type User struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Age       int    `json:"age"`
}

type Trainer struct {
	Name string
	Age  int
	City string
}

func WriteJsonResponseHandler(w http.ResponseWriter, r *http.Request) {
	// 单个插入
	ash := Trainer{"Ash", 10, "Pallet Town"}
	InsertOneResult := models.NewMgo().InsertOne(ash)
	fmt.Println("Inserted a single document: ", InsertOneResult)

	p := User{
		FirstName: "John",
		LastName:  "Doe",
		Age:       25,
	}
	// Set response header
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(&p)
	if err != nil {
		//... handle error
	}
}
