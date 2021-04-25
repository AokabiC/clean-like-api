package infra

import (
	"fmt"
	"go-clean/adapter"
	"go-clean/ent"
	"go-clean/ent/migrate"

	"context"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/lib/pq"
)

func Run() {
	dsn := os.Getenv("DATABASE_URL")
	dbClient, err := ent.Open("postgres", dsn)
	if err != nil {
		fmt.Printf("failed connecting database. %s", err)
	}
	defer dbClient.Close()

	ctx := context.Background()
	err = dbClient.Debug().Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
	if err != nil {
		fmt.Printf("failed creating schema resources. %s", err)
	}

	e := echo.New()

	userController := adapter.NewUserController(dbClient)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/user/:id", func(c echo.Context) error { return userController.GetById(c) })
	e.POST("user/create", func(c echo.Context) error { return userController.Create(c) })

	e.Start("0.0.0.0:8080")
}
