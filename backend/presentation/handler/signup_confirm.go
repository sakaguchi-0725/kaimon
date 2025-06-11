package handler

import (
	"backend/core"
	"backend/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type SignUpConfirmRequest struct {
	Email            string `json:"email"`
	ConfirmationCode string `json:"confirmationCode"`
}

func NewSignUpConfirm(usecase usecase.SignUpConfirm) echo.HandlerFunc {
	return func(c echo.Context) error {
		input, err := makeSignUpConfirmInput(c)
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

func makeSignUpConfirmInput(c echo.Context) (usecase.SignUpConfirmInput, error) {
	var req SignUpConfirmRequest
	if err := c.Bind(&req); err != nil {
		return usecase.SignUpConfirmInput{}, core.NewInvalidError(err)
	}

	return usecase.SignUpConfirmInput{
		Email:            req.Email,
		ConfirmationCode: req.ConfirmationCode,
	}, nil
}
