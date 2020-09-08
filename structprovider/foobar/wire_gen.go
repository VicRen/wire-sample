// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package foobar

// Injectors from wire.go:

func InitializeFooBar() FooBar {
	foo := ProvideFoo()
	bar := ProvideBar()
	fooBar := FooBar{
		MyFoo: foo,
		MyBar: bar,
	}
	return fooBar
}