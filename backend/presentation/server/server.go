package server

import (
	"backend/presentation/handler"
	"backend/presentation/middleware"
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
	s.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
	})

	s.POST("/signup", handler.NewSignUp(uc.SignUp))
	s.POST("/signup/confirm", handler.NewSignUpConfirm(uc.SignUpConfirm))
	s.POST("/login", handler.NewLogin(uc.Login))
	s.POST("/reset-password", handler.NewResetPassword(uc.ResetPassword))
	s.POST("/reset-password/confirm", handler.NewResetPasswordConfirm(uc.ResetPasswordConfirm))
	s.POST("/resend-confirmation-code", handler.NewResendConfirmationCode(uc.ResendConfirmationCode))

	groups := s.Group("/groups")
	groups.Use(middleware.AuthMiddleware(uc.VerifyToken))
	groups.GET("", handler.NewGetJoinedGroups(uc.GetJoinedGroups))
	groups.POST("", handler.NewCreateGroup(uc.CreateGroup))
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
