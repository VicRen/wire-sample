package main

import (
	"flag"
	"github.com/vicren/wire-sample/returnerror/user"
)

var name = flag.String("name", "", "")

func main() {
	flag.Parse()
	u, err := user.NewUser(*name)
	if err != nil {
		panic(err)
	}
	u.Show()
}
