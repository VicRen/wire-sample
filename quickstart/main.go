package main

import (
	"github.com/vicren/wire-sample/quickstart/greeter"
)

func main() {
	event := greeter.InitializeEvent("hello world")

	event.Start()
}
