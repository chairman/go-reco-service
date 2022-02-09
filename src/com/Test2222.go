package main

import "fmt"

func main() {
	arr := [8]int{}
	for i := 0; i < 8; i++ {
		arr[i] = i
	}

	fmt.Println(arr)
	exchangeByAddress(&arr)
	fmt.Println(arr)
}

func exchangeByAddress(arr *[8]int) {
	for k, v := range *arr {
		arr[k] = v * 2
	}
}
