package main

import (
	"fmt"
	"log"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}

func main() {
	fmt.Println("aaaa")
	http.HandleFunc("/", IndexHandler)
	log.Fatal(http.ListenAndServe("0.0.0.0:5000", nil))
}
