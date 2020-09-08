package greeter

import (
	"fmt"
	"github.com/google/wire"
)

type Message struct {
	msg string
}

type Greeter struct {
	Message Message
}

func (g Greeter) Greet() Message {
	return g.Message
}

type Event struct {
	Greeter Greeter
}

func (e Event) Start() {
	msg := e.Greeter.Greet().msg
	fmt.Println(msg)
}

func NewMessage(msg string) Message {
	return Message{
		msg:msg,
	}
}

func NewGreeter(m Message) Greeter {
	return Greeter{Message:m}
}

func NewEvent(g Greeter) Event {
	return Event{
		Greeter:g,
	}
}

var NewSet = wire.NewSet(NewEvent, NewMessage, NewGreeter)