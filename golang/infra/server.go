package infra

import (
	"fmt"
	"go-clean/adapter"
	"go-clean/ent"
	"go-clean/ent/migrate"

	"context"
	"log"
	"os"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
	if err = dbClient.Debug().Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	); err != nil {
		log.Fatalf("failed creating schema resources. %s", err)
	}

	e := echo.New()
	e.Validator = &adapter.RequestValidator{Validator: validator.New()}

	userController := adapter.NewUserController(dbClient)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/users/:id", func(c echo.Context) error { return userController.GetByID(c) })
	e.POST("users/create", func(c echo.Context) error { return userController.Create(c) })

	log.Fatal(e.Start("0.0.0.0:8080"))
}
