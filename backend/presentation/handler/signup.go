package handler

import (
	"backend/core"
	"backend/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type SignUpRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewSignUp(uc usecase.SignUp) echo.HandlerFunc {
	return func(c echo.Context) error {
		input, err := makeSignUpInput(c)
		if err != nil {
			return err
		}

		ctx := c.Request().Context()
		if err := uc.Execute(ctx, input); err != nil {
			return err
		}

		return c.NoContent(http.StatusNoContent)
	}
}

func makeSignUpInput(c echo.Context) (usecase.SignUpInput, error) {
	var req SignUpRequest
	if err := c.Bind(&req); err != nil {
		return usecase.SignUpInput{}, core.NewInvalidError(err)
	}

	return usecase.SignUpInput{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}, nil
}
