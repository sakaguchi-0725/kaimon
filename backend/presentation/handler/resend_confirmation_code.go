package handler

import (
	"backend/core"
	"backend/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ResendConfirmationCodeRequest struct {
	Email string `json:"email" validate:"required,email"`
}

func NewResendConfirmationCode(usecase usecase.ResendConfirmationCode) echo.HandlerFunc {
	return func(c echo.Context) error {
		input, err := makeResendConfirmationCodeInput(c)
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

func makeResendConfirmationCodeInput(c echo.Context) (usecase.ResendConfirmationCodeInput, error) {
	var req ResendConfirmationCodeRequest
	if err := c.Bind(&req); err != nil {
		return usecase.ResendConfirmationCodeInput{}, core.NewInvalidError(err)
	}

	return usecase.ResendConfirmationCodeInput{
		Email: req.Email,
	}, nil
}
