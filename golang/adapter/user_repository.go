package adapter

import (
	"context"
	"go-clean/domain"
	"go-clean/ent"
	entUser "go-clean/ent/user"
)

type UserPgRepository struct {
	client *ent.Client
}

var _ domain.UserRepository = (*UserPgRepository)(nil)

func (repo *UserPgRepository) GetByID(ctx context.Context, uid int) (*domain.User, error) {
	userRecord, err := repo.client.User.Query().Where(entUser.IDEQ(uid)).Only(ctx)
	if err != nil {
		return nil, err
	}

	user := &domain.User{
		ID:   userRecord.ID,
		Name: userRecord.Name,
	}
	return user, nil
}

func (repo *UserPgRepository) Create(ctx context.Context, newUser *domain.User) error {
	userRecord, err := repo.client.User.
		Create().
		SetName(newUser.Name).
		Save(ctx)
	if err != nil {
		return err
	}

	newUser.ID = userRecord.ID
	return nil
}
