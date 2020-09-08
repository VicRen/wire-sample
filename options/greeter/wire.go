// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package greeter

import (
	"context"
	"github.com/google/wire"
	"io"
)

func InitializeGreeter(ctx context.Context, msg []Message, w io.Writer, r io.Reader) (*Greeter, error) {
	wire.Build(Set)
	return nil, nil
}
