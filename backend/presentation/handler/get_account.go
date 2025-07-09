package handler

import (
	"backend/core"
	"backend/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type getAccountResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func NewGetAccount(uc usecase.GetAccount) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		input := usecase.GetAccountInput{
			UserID: core.GetUserID(ctx),
		}

		output, err := uc.Execute(ctx, input)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, getAccountResponse{
			ID:   output.ID,
			Name: output.Name,
		})
	}
}
