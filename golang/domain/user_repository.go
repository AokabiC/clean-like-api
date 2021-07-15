package domain

import (
	"context"
)

type UserRepository interface {
	// IDに一致するUserを返す。見つからない場合はerrを返す。
	GetByID(ctx context.Context, id UserID) (*User, error)
	// Userを新規作成する。
	Create(ctx context.Context, user User) (*User, error)
	// Userを更新する。
	Update(ctx context.Context, user User) (*User, error)
}
