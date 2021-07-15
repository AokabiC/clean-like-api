package domain

import (
	"context"
)

type UserRepository interface {
	GetByID(ctx context.Context, id UserID) (*User, error)
	Create(ctx context.Context, user User) (*User, error)
}
