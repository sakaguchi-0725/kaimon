package handler

import (
	"backend/core"
	"backend/usecase"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

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
	authHeader := c.Request().Header.Get("Authorization")
	if authHeader == "" {
		return usecase.SignUpInput{}, core.NewInvalidError(errors.New("authorization header is required"))
	}

	// Bearer tokenのプレフィックスを除去
	idToken, err := core.RemovePrefix(authHeader, "Bearer ")
	if err != nil {
		return usecase.SignUpInput{}, core.NewInvalidError(err)
	}

	return usecase.SignUpInput{
		IDToken: idToken,
	}, nil
}
