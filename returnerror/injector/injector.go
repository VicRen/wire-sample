//+build wireinject

package injector

import (
	"github.com/google/wire"
	"github.com/vicren/wire-sample/returnerror/user"
)

func InitializeUser(name string) (*user.User, error) {
	wire.Build(user.NewUser)
	return nil, nil
}