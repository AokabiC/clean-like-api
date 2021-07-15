package adapter

import (
	"go-clean/domain"
	"go-clean/ent"
	"go-clean/usecase"

	"github.com/labstack/echo/v4"

	"context"
	"errors"
	"net/http"
	"strconv"
)

type UserController struct {
	Interactor usecase.UserUsecase
}

func NewUserController(client *ent.Client) *UserController {
	return &UserController{
		Interactor: &usecase.UserInteractor{
			UserRepository: &UserPgRepository{
				client: client,
			},
		},
	}
}

type (
	// GET users/:id
	UserGetByIDResponse struct {
		ID       int    `json:"id"`
		Username string `json:"username"`
	}
)

func (controller *UserController) GetByID(c echo.Context) (err error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	ctx := context.Background()
	user, err := controller.Interactor.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	response := &UserGetByIDResponse{
		ID:       int(user.ID),
		Username: string(user.Username),
	}

	c.JSON(http.StatusOK, response)
	return nil
}

type (
	// POST users/create
	UserCreateRequest struct {
		Username string `json:"username" validate:"required"`
	}
	UserCreateResponse struct {
		ID       int    `json:"id"`
		Username string `json:"username"`
	}
)

func (controller *UserController) Create(c echo.Context) (err error) {
	request := new(UserCreateRequest)
	if err := c.Bind(request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	if err := c.Validate(request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	ctx := context.Background()

	user, err := controller.Interactor.Create(ctx, request.Username)
	if err != nil {
		c.Echo().Logger.Error(err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	response := &UserCreateResponse{
		ID:       int(user.ID),
		Username: string(user.Username),
	}
	_ = c.JSON(http.StatusCreated, response)
	return nil
}

type (
	// POST users/update_username
	UserUpdateUsernameRequest struct {
		ID       int    `json:"id" validate:"required"`
		Username string `json:"username" validate:"required"`
	}
)

func (controller *UserController) UpdateUsername(c echo.Context) (err error) {
	request := new(UserUpdateUsernameRequest)
	if err := c.Bind(request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	if err := c.Validate(request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	ctx := context.Background()

	_, err = controller.Interactor.UpdateUsername(ctx, request.ID, request.Username)
	if err != nil {
		c.Echo().Logger.Error(err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return echo.NewHTTPError(http.StatusOK)
}
