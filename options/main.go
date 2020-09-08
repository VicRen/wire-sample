package main

import (
	"context"
	"github.com/vicren/wire-sample/options/greeter"
)

func main() {
	g, err := greeter.InitializeGreeter(context.Background(), []greeter.Message{"1","2"}, nil, nil)
	if err != nil {
		panic(err)
	}
	g.ShowMessage()
}
