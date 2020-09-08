package main

import (
	"fmt"
	"github.com/vicren/wire-sample/structprovider/foobar"
)

func main() {
	fb := foobar.InitializeFooBar()

	fmt.Println(fb.MyFoo, fb.MyBar)
}
