package user

import (
	"errors"
	"fmt"
)

type Repo struct {
	User string
}

func (r*Repo) Show() {
	fmt.Println(r.User)
}

func NewRepo(name string) (*Repo, func(), error) {
	if name == "" {
		return nil, nil, errors.New("name cannot be empty")
	}
	f := func() {
		fmt.Println("Now showing")
	}
	return &Repo{User:name}, f, nil
}

