package user

import (
	"errors"
	"fmt"
)

type User struct {
	Name string
}

func (u *User) Show() {
	fmt.Println(u.Name)
}

func NewUser(name string) (*User, error) {
	if name == "" {
		return nil, errors.New("name cannot be empty")
	}
	return &User{name}, nil
}