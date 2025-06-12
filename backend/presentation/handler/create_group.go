package handler

import (
	"backend/core"
	"backend/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CreateGroupRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func NewCreateGroup(usecase usecase.CreateGroup) echo.HandlerFunc {
	return func(c echo.Context) error {
		input, err := makeCreateGroupInput(c)
		if err != nil {
			return err
		}

		ctx := c.Request().Context()
		if err := usecase.Execute(ctx, input); err != nil {
			return err
		}

		return c.NoContent(http.StatusCreated)
	}
}

func makeCreateGroupInput(c echo.Context) (usecase.CreateGroupInput, error) {
	var req CreateGroupRequest
	if err := c.Bind(&req); err != nil {
		return usecase.CreateGroupInput{}, core.NewInvalidError(err)
	}

	return usecase.CreateGroupInput{
		UserID:      core.GetUserID(c.Request().Context()),
		Name:        req.Name,
		Description: req.Description,
	}, nil
}
