//+build wireinject

package injector

import (
	"github.com/google/wire"
	"github.com/vicren/wire-sample/returnfunc/user"
)

func InitializeUser(name string) (*user.Repo, func(), error) {
	wire.Build(user.NewRepo)
	return nil, nil, nil
}