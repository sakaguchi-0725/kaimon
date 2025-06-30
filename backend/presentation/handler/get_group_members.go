package handler

import (
	"backend/core"
	"backend/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type GetGroupMembersResponse struct {
	Members []Member `json:"members"`
}

type Member struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Role   string `json:"role"`
	Status string `json:"status"`
}

func NewGetGroupMembers(usecase usecase.GetGroupMembers) echo.HandlerFunc {
	return func(c echo.Context) error {
		outputs, err := usecase.Execute(c.Request().Context(), makeGetGroupMembersInput(c))
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, makeGetGroupMembersResponse(outputs))
	}
}

func makeGetGroupMembersInput(c echo.Context) usecase.GetGroupMembersInput {
	return usecase.GetGroupMembersInput{
		GroupID: c.Param("id"),
		UserID:  core.GetUserID(c.Request().Context()),
	}
}

func makeGetGroupMembersResponse(outputs usecase.GetGroupMembersOutput) GetGroupMembersResponse {
	members := make([]Member, len(outputs.Members))
	for i, member := range outputs.Members {
		members[i] = Member{
			ID:     member.ID,
			Name:   member.Name,
			Role:   member.Role,
			Status: member.Status,
		}
	}
	return GetGroupMembersResponse{
		Members: members,
	}
}
