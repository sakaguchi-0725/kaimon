package handler

import (
	"backend/core"
	"backend/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type groupInvitationResponse struct {
	Code      string `json:"code"`
	ExpiresAt string `json:"expiresAt"`
}

func NewGroupInvitation(uc usecase.GroupInvitation) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		input := usecase.GroupInvitationInput{
			GroupID: c.Param("id"),
			UserID:  core.GetUserID(ctx),
		}

		output, err := uc.Execute(ctx, input)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, groupInvitationResponse{
			Code:      output.Code,
			ExpiresAt: output.ExpiresAt,
		})
	}
}
