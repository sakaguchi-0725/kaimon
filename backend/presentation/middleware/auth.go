package middleware

import (
	"backend/core"
	"backend/usecase"
	"context"

	"github.com/labstack/echo/v4"
)

func AuthMiddleware(usecase usecase.VerifyToken) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := c.Request().Context()
			authHeader := c.Request().Header.Get("Authorization")

			token, err := core.RemovePrefix(authHeader, "Bearer ")
			if err != nil {
				return core.NewInvalidError(err)
			}

			userID, email, err := usecase.Execute(ctx, token)
			if err != nil {
				return err
			}

			newCtx := context.WithValue(ctx, core.UserIDKey, userID)
			newCtx = context.WithValue(newCtx, core.EmailKey, email)
			c.SetRequest(c.Request().WithContext(newCtx))

			return next(c)
		}
	}
}
