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

			// TODO: 認証機能を実装したら有効化
			// cookie, err := c.Cookie("access_token")
			// if err != nil {
			// 	if err == http.ErrNoCookie {
			// 		return core.NewAppError(core.ErrUnauthorized, err)
			// 	}

			// 	return err
			// }
			userID, err := usecase.Execute(ctx, "dummy-token")
			if err != nil {
				return err
			}

			newCtx := context.WithValue(ctx, core.UserIDKey, userID)
			c.SetRequest(c.Request().WithContext(newCtx))

			return next(c)
		}
	}
}
