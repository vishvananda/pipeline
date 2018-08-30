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

func apply(requestContext *gp.RequestContext, ignore string, in interface{}) interface{} {
	return in
}
