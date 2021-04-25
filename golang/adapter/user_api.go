package adapter

import (
	"go-clean/ent"
	"go-clean/usecase"

	"github.com/labstack/echo"

	"context"
	"strconv"
)

type UserController struct {
	Interactor usecase.UserInteractor
}

func NewUserController(client *ent.Client) *UserController {
	return &UserController{
		Interactor: usecase.UserInteractor{
			UserRepository: &UserPgRepository{
				client: client,
			},
		},
	}
}

func (controller *UserController) GetById(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	ctx := context.Background()
	user, err := controller.Interactor.GetById(ctx, id)
	if err != nil {
		if ent.IsNotFound(err) {
			c.JSON(404, "")
			return
		}
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, user)
	return
}

func (controller *UserController) Create(c echo.Context) (err error) {
	item := new(UserCreateRequest)
	if err := c.Bind(item); err != nil {
		return err
	}

	ctx := context.Background()

	user, err := controller.Interactor.Create(ctx, item.Name)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, user)
	return
}

type UserCreateRequest struct {
	Name string `json:"name"`
}
