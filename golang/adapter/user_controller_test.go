package adapter

import (
	"go-clean/domain"
	mock "go-clean/usecase/mock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-playground/validator"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func prepareUser(username string) *domain.User {
	uname, _ := domain.NewUsername(username)
	user, _ := domain.NewUser(
		1,
		uname,
	)
	return user
}

func TestUserController_GetByID(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name          string
		args          args
		prepareMockFn func(*mock.MockUserUsecase)
		want          string
		wantErr       bool
	}{
		{
			name: "OK",
			args: args{
				id: "1",
			},
			prepareMockFn: func(m *mock.MockUserUsecase) {
				user := prepareUser("test_user1")
				m.EXPECT().GetByID(gomock.Any(), 1).Return(user, nil)
			},
			want: `{
				"id": 1,
				"username":"test_user1"
			}`,
			wantErr: false,
		},
		{
			name: "NotFound",
			args: args{
				id: "404",
			},
			prepareMockFn: func(m *mock.MockUserUsecase) {
				m.EXPECT().GetByID(gomock.Any(), 404).Return(nil, domain.ErrUserNotFound)
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "BadRequestContainAlphabet",
			args: args{
				id: "1e",
			},
			prepareMockFn: func(m *mock.MockUserUsecase) {
				// 呼ばれてはならない
				m.EXPECT().GetByID(gomock.Any(), gomock.Any()).Times(0)
			},
			wantErr: true,
		},
		{
			name: "BadRequestTooLongID",
			args: args{
				id: "10000000000000000000000000",
			},
			prepareMockFn: func(m *mock.MockUserUsecase) {
				// 呼ばれてはならない
				m.EXPECT().GetByID(gomock.Any(), gomock.Any()).Times(0)
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// create mock
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockUserInteractor := mock.NewMockUserUsecase(ctrl)
			tt.prepareMockFn(mockUserInteractor)
			// DI
			controller := &UserController{mockUserInteractor}

			e := echo.New()
			e.Validator = &RequestValidator{Validator: validator.New()}
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/users/:id")
			c.SetParamNames("id")
			c.SetParamValues(tt.args.id)

			err := controller.GetByID(c)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserController.GetByID() error = %v, wantErr %v", err, tt.wantErr)
			}
			if (err == nil) && !assert.JSONEq(t, rec.Body.String(), tt.want) {
				t.Errorf("UserController.GetByID() = %v, want %v", rec.Body.String(), tt.want)
			}
		})
	}

}

func TestUserController_Create(t *testing.T) {
	type args struct {
		body string
	}
	tests := []struct {
		name          string
		args          args
		prepareMockFn func(*mock.MockUserUsecase)
		want          string
		wantErr       bool
	}{
		{
			name: "OK",
			args: args{
				body: `{
					"username":"test_user1"
				}`,
			},
			prepareMockFn: func(m *mock.MockUserUsecase) {
				user := prepareUser("test_user1")
				m.EXPECT().Create(
					gomock.Any(),
					"test_user1",
				).Return(user, nil)
			},
			want: `{
				"id": 1,
				"username":"test_user1"
			}`,
			wantErr: false,
		},
		{
			name: "BadRequest",
			args: args{
				body: `{}`,
			},
			prepareMockFn: func(m *mock.MockUserUsecase) {
				// 呼ばれてはならない
				m.EXPECT().Create(gomock.Any(), gomock.Any()).Times(0)
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// create mock
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockUserInteractor := mock.NewMockUserUsecase(ctrl)
			tt.prepareMockFn(mockUserInteractor)
			// DI
			controller := &UserController{mockUserInteractor}

			e := echo.New()
			e.Validator = &RequestValidator{Validator: validator.New()}
			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(tt.args.body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/users/create")

			err := controller.Create(c)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserController.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
			if (err == nil) && !assert.JSONEq(t, rec.Body.String(), tt.want) {
				t.Errorf("UserController.Create() = %v, want %v", rec.Body.String(), tt.want)
			}
		})
	}
}

func TestUserController_UpdateUsername(t *testing.T) {
	type args struct {
		body string
	}
	tests := []struct {
		name          string
		args          args
		prepareMockFn func(*mock.MockUserUsecase)
		wantErr       bool
	}{
		{
			name: "OK",
			args: args{
				body: `{
					"id": 1,
					"username":"test_user1"
				}`,
			},
			prepareMockFn: func(m *mock.MockUserUsecase) {
				user := prepareUser("test_user1")
				m.EXPECT().UpdateUsername(
					gomock.Any(),
					1,
					"test_user1",
				).Return(user, nil)
			},
			wantErr: false,
		},
		{
			name: "BadRequest",
			args: args{
				body: `{}`,
			},
			prepareMockFn: func(m *mock.MockUserUsecase) {
				// 呼ばれてはならない
				m.EXPECT().UpdateUsername(gomock.Any(), gomock.Any(), gomock.Any()).Times(0)
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// create mock
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockUserInteractor := mock.NewMockUserUsecase(ctrl)
			tt.prepareMockFn(mockUserInteractor)
			// DI
			controller := &UserController{mockUserInteractor}

			e := echo.New()
			e.Validator = &RequestValidator{Validator: validator.New()}
			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(tt.args.body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/users/update_username")

			err := controller.UpdateUsername(c)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserController.UpdateUsername() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
