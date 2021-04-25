package domain

import (
	"context"
	"fmt"
)

type User struct {
	ID   int
	Name string
}

func NewUser(name string) (*User, error) {
	if len(name) < 2 {
		return nil, fmt.Errorf("ユーザー名は2文字以上です")
	}

	return &User{
		ID:   0,
		Name: name,
	}, nil
}

type UserRepository interface {
	GetByID(ctx context.Context, uid int) (user *User, err error)
	Create(ctx context.Context, newUser *User) (err error)
}
