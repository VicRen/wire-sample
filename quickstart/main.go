package main

import (
	"github.com/vicren/wire-sample/quickstart/greeter"
)

func main() {
	event := greeter.InitializeEvent("hello_world")

	event.Start()
}
