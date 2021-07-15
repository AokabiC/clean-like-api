package usecase

import (
	"context"
	"fmt"
	"go-clean/domain"
)

type UserUsecase interface {
	// IDに一致するUserを返す。見つからない場合はerrを返す。
	GetByID(ctx context.Context, id int) (*domain.User, error)
	// 与えられた名前でUserを作成し、作成されたUserを返す。
	Create(ctx context.Context, username string) (*domain.User, error)
}

type UserInteractor struct {
	UserRepository domain.UserRepository
}

var _ UserUsecase = (*UserInteractor)(nil)

func (interactor *UserInteractor) GetByID(ctx context.Context, id int) (*domain.User, error) {
	user, err := interactor.UserRepository.GetByID(ctx, domain.UserID(id))
	return user, err
}

func (interactor *UserInteractor) Create(ctx context.Context, username string) (*domain.User, error) {
	newUsername, err := domain.NewUsername(username)
	if err != nil {
		return nil, ErrInvalidUserCreateRequest(err)
	}

	userWithoutID, err := domain.NewUser(-1, newUsername)
	if err != nil {
		return nil, err
	}

	user, err := interactor.UserRepository.Create(ctx, *userWithoutID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func ErrInvalidUserCreateRequest(err error) error {
	return fmt.Errorf("invalid user create request: %w", err)
}
