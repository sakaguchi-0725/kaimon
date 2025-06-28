package handler

import (
	"backend/core"
	"backend/usecase"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

type SignUpRequest struct {
	Name            string `json:"name"`
	ProfileImageURL string `json:"profileImageUrl"`
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

	ctx := c.Request().Context()
	uid := core.GetUserID(ctx)
	if uid == "" {
		return usecase.SignUpInput{}, core.NewInvalidError(errors.New("userID is required"))
	}

	email := core.GetEmail(ctx)
	if email == "" {
		return usecase.SignUpInput{}, core.NewInvalidError(errors.New("email is required"))
	}

	return usecase.SignUpInput{
		UID:             uid,
		Email:           email,
		Name:            req.Name,
		ProfileImageURL: req.ProfileImageURL,
	}, nil
}
