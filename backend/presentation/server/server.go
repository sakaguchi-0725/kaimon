package server

import (
	"backend/core"
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

func (s *Server) MapRoutes(frontendURL string, logger core.Logger, uc registry.UseCase) {
	api := s.Group("/api")

	api.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{frontendURL},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	api.Use(middleware.Logger())
	api.Use(customMiddleware.AuthMiddleware(uc.VerifyToken))
	api.Use(customMiddleware.Error(logger))

	api.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
	})

	api.POST("/signup", handler.NewSignUp(uc.SignUp))

	groups := api.Group("/groups")
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
