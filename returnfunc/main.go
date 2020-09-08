package  main

import (
	"flag"
	"github.com/vicren/wire-sample/returnfunc/injector"
)

var name = flag.String("name", "", "")

func main() {
	flag.Parse()
	r, f, err := injector.InitializeUser(*name)
	if err != nil {
		panic(err)
	}
	f()
	r.Show()
}
