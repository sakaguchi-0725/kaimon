package handler

import (
	"backend/core"
	"backend/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewLogin(usecase usecase.Login) echo.HandlerFunc {
	return func(c echo.Context) error {
		input, err := makeLoginInput(c)
		if err != nil {
			return err
		}

		ctx := c.Request().Context()
		if err := usecase.Execute(ctx, input); err != nil {
			return err
		}

		return c.NoContent(http.StatusNoContent)
	}
}

func makeLoginInput(c echo.Context) (usecase.LoginInput, error) {
	var req LoginRequest
	if err := c.Bind(&req); err != nil {
		return usecase.LoginInput{}, core.NewInvalidError(err)
	}

	return usecase.LoginInput{
		Email:    req.Email,
		Password: req.Password,
	}, nil
}
