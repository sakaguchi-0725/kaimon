package server

import (
	"backend/presentation/handler"
	customMiddleware "backend/presentation/middleware"
	"backend/registry"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	addr string
	*echo.Echo
}

func New(addr uint16) *Server {
	return &Server{
		addr: fmt.Sprintf(":%d", addr),
		Echo: echo.New(),
	}
}

func (s *Server) MapRoutes(uc registry.UseCase) {
	s.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:5173"}, // TODO: 環境変数から取得する
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowCredentials: true,
	}))

	s.Use(middleware.Logger())

	s.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
	})

	s.POST("/signup", handler.NewSignUp(uc.SignUp))
	s.POST("/reset-password", handler.NewResetPassword(uc.ResetPassword))
	s.POST("/reset-password/confirm", handler.NewResetPasswordConfirm(uc.ResetPasswordConfirm))
	s.POST("/resend-confirmation-code", handler.NewResendConfirmationCode(uc.ResendConfirmationCode))

	groups := s.Group("/groups")
	groups.Use(customMiddleware.AuthMiddleware(uc.VerifyToken))
	groups.GET("", handler.NewGetJoinedGroups(uc.GetJoinedGroups))
	groups.POST("", handler.NewCreateGroup(uc.CreateGroup))
	groups.GET("/:id", handler.NewGetGroup(uc.GetGroup))
	groups.GET("/:id/members", handler.NewGetGroupMembers(uc.GetGroupMembers))
	groups.GET("/:id/items", handler.NewGetShoppingItems(uc.GetShoppingItems))
}

func (s *Server) Run() {
	go func() {
		if err := s.Start(s.addr); err != nil && err != http.ErrServerClosed {
			log.Fatalf("failed to start server: %v\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		log.Fatalf("failed to shutdown server: %v\n", err)
	}

	log.Println("server shutdown successfully")
}
