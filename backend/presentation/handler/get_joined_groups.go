package handler

import (
	"backend/core"
	"backend/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	GetJoinedGroupsResponse struct {
		Groups []JoinedGroup `json:"groups"`
	}

	JoinedGroup struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
	}
)

func NewGetJoinedGroups(usecase usecase.GetJoinedGroups) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		groups, err := usecase.Execute(ctx, core.GetUserID(ctx))
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, makeGetJoinedGroupsResponse(groups))
	}
}

func makeGetJoinedGroupsResponse(outputs []usecase.GetJoinedGroupOutput) GetJoinedGroupsResponse {
	groups := make([]JoinedGroup, len(outputs))
	for i, output := range outputs {
		groups[i] = JoinedGroup{
			ID:          output.ID,
			Name:        output.Name,
			Description: output.Description,
		}
	}

	return GetJoinedGroupsResponse{
		Groups: groups,
	}
}
