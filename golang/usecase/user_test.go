package usecase

import (
	"context"
	"go-clean/domain"
	mock "go-clean/domain/mock"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
)

func prepareUser(username string) *domain.User {
	uname, _ := domain.NewUsername(username)
	user, _ := domain.NewUser(
		1,
		uname,
	)
	return user
}

func TestUserInteractor_GetByID(t *testing.T) {
	type args struct {
		ctx context.Context
		id  int
	}
	tests := []struct {
		name          string
		args          args
		prepareMockFn func(*mock.MockUserRepository)
		want          *domain.User
		wantErr       bool
	}{
		{
			name: "OK",
			args: args{
				ctx: context.Background(),
				id:  1,
			},
			prepareMockFn: func(m *mock.MockUserRepository) {
				m.EXPECT().GetByID(gomock.Any(), domain.UserID(1)).Return(prepareUser("test_user1"), nil)
			},
			want: &domain.User{
				ID:       domain.UserID(1),
				Username: "test_user1",
			},
			wantErr: false,
		},
		{
			name: "NotFound",
			args: args{
				ctx: context.Background(),
				id:  404,
			},
			prepareMockFn: func(m *mock.MockUserRepository) {
				m.EXPECT().GetByID(gomock.Any(), domain.UserID(404)).Return(nil, domain.ErrUserNotFound)
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// create mock
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockUserRepository := mock.NewMockUserRepository(ctrl)
			tt.prepareMockFn(mockUserRepository)
			// DI
			interactor := &UserInteractor{mockUserRepository}

			got, err := interactor.GetByID(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserInteractor.GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserInteractor.GetByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserInteractor_Create(t *testing.T) {
	type args struct {
		ctx      context.Context
		username string
	}
	tests := []struct {
		name          string
		args          args
		prepareMockFn func(*mock.MockUserRepository)
		want          *domain.User
		wantErr       bool
	}{
		{
			name: "OK",
			args: args{
				ctx:      context.Background(),
				username: "test_user1",
			},
			prepareMockFn: func(m *mock.MockUserRepository) {
				userWithoutID, _ := domain.NewUser(-1, "test_user1")
				m.EXPECT().Create(gomock.Any(), *userWithoutID).Return(prepareUser("test_user1"), nil)
			},
			want: &domain.User{
				ID:       domain.UserID(1),
				Username: "test_user1",
			},
			wantErr: false,
		},
		{
			name: "InvalidUsername",
			args: args{
				ctx:      context.Background(),
				username: "test_too______long______name",
			},
			prepareMockFn: func(m *mock.MockUserRepository) {
				// 呼ばれてはならない
				m.EXPECT().Create(gomock.Any(), gomock.Any()).Times(0)
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// create mock
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockUserRepository := mock.NewMockUserRepository(ctrl)
			tt.prepareMockFn(mockUserRepository)
			// DI
			interactor := &UserInteractor{mockUserRepository}

			got, err := interactor.Create(tt.args.ctx, tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserInteractor.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserInteractor.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserInteractor_UpdateUsername(t *testing.T) {
	type args struct {
		ctx      context.Context
		id       int
		username string
	}
	tests := []struct {
		name          string
		args          args
		prepareMockFn func(*mock.MockUserRepository)
		want          *domain.User
		wantErr       bool
	}{
		{
			name: "OK",
			args: args{
				ctx:      context.Background(),
				id:       1,
				username: "updated_user1",
			},
			prepareMockFn: func(m *mock.MockUserRepository) {
				user := prepareUser("test_user1")
				m.EXPECT().GetByID(gomock.Any(), domain.UserID(1)).Return(user, nil)
				user.Username = "updated_user1"
				m.EXPECT().Update(gomock.Any(), *user).Return(user, nil)
			},
			want: &domain.User{
				ID:       domain.UserID(1),
				Username: "updated_user1",
			},
			wantErr: false,
		},
		{
			name: "InvalidUsername",
			args: args{
				ctx:      context.Background(),
				id:       1,
				username: "test_too______long______name",
			},
			prepareMockFn: func(m *mock.MockUserRepository) {
				// 呼ばれてはならない
				m.EXPECT().GetByID(gomock.Any(), gomock.Any()).Times(0)
				m.EXPECT().Update(gomock.Any(), gomock.Any()).Times(0)
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// create mock
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockUserRepository := mock.NewMockUserRepository(ctrl)
			tt.prepareMockFn(mockUserRepository)
			// DI
			interactor := &UserInteractor{mockUserRepository}
			got, err := interactor.UpdateUsername(tt.args.ctx, tt.args.id, tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserInteractor.UpdateUsername() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserInteractor.UpdateUsername() = %v, want %v", got, tt.want)
			}
		})
	}
}
