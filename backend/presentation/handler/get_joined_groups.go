package handler

import (
	"backend/core"
	"backend/usecase"
	"errors"
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
		MemberCount int    `json:"memberCount"`
	}
)

func NewGetJoinedGroups(usecase usecase.GetJoinedGroups) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		userID := core.GetUserID(ctx)
		if userID == "" {
			return core.NewAppError(core.ErrUnauthorized, errors.New("unauthorized"))
		}

		groups, err := usecase.Execute(ctx, userID)
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
			MemberCount: output.MemberCount,
		}
	}

	return GetJoinedGroupsResponse{
		Groups: groups,
	}
}
