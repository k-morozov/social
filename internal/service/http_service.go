package service

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
)

type ServiceHTTP struct {
	server *http.Server
	engine *echo.Echo
	logger zerolog.Logger
}

type ServiceHTTPOption func(s *ServiceHTTP)

func NewServiceHTTP(opts ...ServiceHTTPOption) (*ServiceHTTP, error) {
	srv := &ServiceHTTP{
		engine: echo.New(),
	}

	for _, opt := range opts {
		opt(srv)
	}

	return srv, nil
}

func (s *ServiceHTTP) ListenAndServe() error {

	s.engine.POST("/user/register", register)
	s.server = &http.Server{
		Addr:    ":8080",
		Handler: s.engine,
	}

	return s.server.ListenAndServe()
}

func register(c echo.Context) error {
	fmt.Println("Call register")

	return c.String(http.StatusOK, "answer from register")
}
