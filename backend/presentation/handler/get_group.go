package handler

import (
	"backend/core"
	"backend/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type GetGroupResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   string `json:"createdAt"`
}

func NewGetGroup(usecase usecase.GetGroup) echo.HandlerFunc {
	return func(c echo.Context) error {
		group, err := usecase.Execute(c.Request().Context(), makeGetGroupInput(c))
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, makeGetGroupResponse(group))
	}
}

func makeGetGroupInput(c echo.Context) usecase.GetGroupInput {
	return usecase.GetGroupInput{
		UserID:  core.GetUserID(c.Request().Context()),
		GroupID: c.Param("id"),
	}
}

func makeGetGroupResponse(group usecase.GetGroupOutput) GetGroupResponse {
	return GetGroupResponse{
		ID:          group.ID,
		Name:        group.Name,
		Description: group.Description,
		CreatedAt:   group.CreatedAt,
	}
}
