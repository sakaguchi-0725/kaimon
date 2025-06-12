package handler

import (
	"backend/core"
	"backend/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ResetPasswordRequest struct {
	Email string `json:"email"`
}

func NewResetPassword(usecase usecase.ResetPassword) echo.HandlerFunc {
	return func(c echo.Context) error {
		input, err := makeResetPasswordInput(c)
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

func makeResetPasswordInput(c echo.Context) (usecase.ResetPasswordInput, error) {
	var req ResetPasswordRequest
	if err := c.Bind(&req); err != nil {
		return usecase.ResetPasswordInput{}, core.NewInvalidError(err)
	}

	return usecase.ResetPasswordInput{
		Email: req.Email,
	}, nil
}
