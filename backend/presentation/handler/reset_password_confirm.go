package handler

import (
	"backend/core"
	"backend/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ResetPasswordConfirmRequest struct {
	Email            string `json:"email" validate:"required,email"`
	ConfirmationCode string `json:"confirmationCode" validate:"required"`
}

func NewResetPasswordConfirm(usecase usecase.ResetPasswordConfirm) echo.HandlerFunc {
	return func(c echo.Context) error {
		input, err := makeResetPasswordConfirmInput(c)
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

func makeResetPasswordConfirmInput(c echo.Context) (usecase.ResetPasswordConfirmInput, error) {
	var req ResetPasswordConfirmRequest
	if err := c.Bind(&req); err != nil {
		return usecase.ResetPasswordConfirmInput{}, core.NewInvalidError(err)
	}

	return usecase.ResetPasswordConfirmInput{
		Email:            req.Email,
		ConfirmationCode: req.ConfirmationCode,
	}, nil
}
