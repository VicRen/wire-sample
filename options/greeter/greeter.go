package greeter

import (
	"context"
	"fmt"
	"github.com/google/wire"
	"io"
)

type Message string

type Options struct {
	Message []Message
	Writer io.Writer
	Reader io.Reader
}

type Greeter struct {
	Messages []Message
}

func (g *Greeter) ShowMessage() {
	for _, m := range g.Messages {
		fmt.Println(m)
	}
}

func NewGreeter(ctx context.Context, opts *Options) (*Greeter, error) {
	return &Greeter{Messages:opts.Message}, nil
}

var Set = wire.NewSet(wire.Struct(new(Options), "*"), NewGreeter)