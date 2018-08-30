package main

import (
	"fmt"

	gp "github.com/oracle/graphpipe-go"
)

func main() {
	model1 := "http://127.0.0.1:3000"
	model2 := "http://127.0.0.1:4000"
	request := [][]float32{{0.0, 1.0}, {2.0, 3.0}}
	fmt.Println(request)
	result1, err := gp.Remote(model1, request)
	if err != nil {
		panic(err)
	}
	fmt.Println(result1)
	result2, err := gp.Remote(model2, result1)
	if err != nil {
		panic(err)
	}
	fmt.Println(result2)
}
