package adapter

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"net/http"
)

type RequestValidator struct {
	Validator *validator.Validate
}

func (rv *RequestValidator) Validate(i interface{}) error {
	if err := rv.Validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return nil
}
