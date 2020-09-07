package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/google/wire"
)

type User struct {
}

type UserService struct {
	UserRepo UserRepository
}

func NewUserService(repository UserRepository) *UserService {
	return &UserService{
		UserRepo:repository,
	}
}

func (u *UserService) UserExist(id int) bool {
	_, err := u.UserRepo.GetUserByID(id)
	return err == nil
}

type UserRepository interface {
	GetUserByID(id int) (*User, error)
}

type mockUserRepo struct {
	foo string
	bar int
}

func NewMockUserRepo(foo string, bar int) *mockUserRepo {
	return &mockUserRepo{
		foo:foo,
		bar:bar,
	}
}

func (u *mockUserRepo) GetUserByID(id int) (*User, error) {
	if id == 0 {
		return nil, errors.New("not found")
	}
	return &User{}, nil
}

var MockUserRepoSet = wire.NewSet(NewMockUserRepo, wire.Bind(new(UserRepository), new(*mockUserRepo)))

var id = flag.Int("id", 0, "")

func main() {
	flag.Parse()
	u := InitializeUserService("foo", 0)
	fmt.Println(u.UserExist(*id))
}