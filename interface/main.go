package main

import (
	"flag"
	"fmt"
	"github.com/vicren/wire-sample/interface/injector"
)

var id = flag.Int("id", 0, "")

func main() {
	flag.Parse()
	u := injector.InitializeUserService("foo", 0)
	fmt.Println(u.UserExist(*id))
}