package usecase

import (
	"context"
	"go-clean/domain"
)

type UserInteractor struct {
	UserRepository domain.UserRepository
}

func (interactor *UserInteractor) GetById(ctx context.Context, id int) (*domain.User, error) {
	user, err := interactor.UserRepository.GetByID(ctx, id)
	return user, err
}

func (interactor *UserInteractor) Create(ctx context.Context, name string) (*domain.User, error) {
	user, err := domain.NewUser(name)
	if err != nil {
		return nil, err
	}
	err = interactor.UserRepository.Create(ctx, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
