package main

import (
	"os"

	gp "github.com/oracle/graphpipe-go"
)

func main() {
	port := "9000"
	if len(os.Args) > 1 {
		port = os.Args[1]
	}
	if err := gp.Serve("0.0.0.0:"+port, false, apply, nil, nil); err != nil {
		panic(err)
	}
}

const (
	nextModelUri = "http://127.0.0.1:4000"
)

func apply(requestContext *gp.RequestContext, _config string, in [][]float32) [][]float32 {
	// multiply every value by 2.0
	out := make([][]float32, len(in))
	for i := range in {
		out[i] = make([]float32, len(in[i]))
		for j := range in[i] {
			out[i][j] = in[i][j] * 2.0
		}
	}

	// make the remote request to multiply
	result, err := gp.Remote(nextModelUri, out)
	if err != nil {
		panic(err)
	}
	// return the result
	return result.([][]float32)
}
