package main

import (
	"fmt"
	"os"

	gp "github.com/oracle/graphpipe-go"
)

func main() {
	port := "3000"
	if len(os.Args) > 1 {
		port = os.Args[1]
	}
	uri := "http://127.0.0.1:" + port
	request := [][]float32{{0.0, 1.0}, {2.0, 3.0}}
	result, err := gp.Remote(uri, request)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
