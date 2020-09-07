// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package injector

import (
	"github.com/google/wire"
	"github.com/vicren/wire-sample/interface/user"
)

func InitializeUserService(foo string, bar int) *user.UserService {
	wire.Build(user.NewUserService, user.MockUserRepoSet)
	return nil
}