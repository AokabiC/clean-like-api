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

func (repo *UserPgRepository) GetByID(ctx context.Context, id domain.UserID) (*domain.User, error) {
	userRecord, err := repo.client.User.Query().Where(entUser.IDEQ(int(id))).Only(ctx)
	if err != nil {
		return nil, err
	}
	user, _ := domain.NewUser(
		domain.UserID(userRecord.ID),
		domain.Username(userRecord.Username),
	)
	return user, nil
}

func (repo *UserPgRepository) Create(ctx context.Context, userWithoutID domain.User) (*domain.User, error) {
	userRecord, err := repo.client.User.
		Create().
		SetUsername(string(userWithoutID.Username)).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	user := &userWithoutID
	user.ID = domain.UserID(userRecord.ID)
	return user, nil
}

func (repo *UserPgRepository) Update(ctx context.Context, user domain.User) (*domain.User, error) {
	userRecord, err := repo.client.User.Query().Where(entUser.IDEQ(int(user.ID))).Only(ctx)
	if err != nil {
		return nil, err
	}

	updatedUserRecord, err := userRecord.
		Update().
		SetUsername(string(user.Username)).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	updatedUser, _ := domain.NewUser(
		domain.UserID(updatedUserRecord.ID),
		domain.Username(updatedUserRecord.Username),
	)

	return updatedUser, nil
}
