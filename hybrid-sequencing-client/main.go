package main

import (
	"fmt"
	"net/http"

	gp "github.com/oracle/graphpipe-go"
)

func main() {
	uri := "http://127.0.0.1:3000"
	request := [][]float32{{0.0, 1.0}, {2.0, 3.0}}
	config := "http://127.0.0.1:4000"
	in := []interface{}{request}
	results, err := gp.MultiRemote(http.DefaultClient, uri, config, in, nil, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(results[0])
}
