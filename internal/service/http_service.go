package service

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"

	"github.com/k-morozov/social/internal/api"
)

type ServiceHTTP struct {
	server *http.Server
	engine *echo.Echo
	logger zerolog.Logger
	api    api.ServerAPI
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
	s.server = &http.Server{
		Addr:    ":8080",
		Handler: s.engine,
	}

	return s.server.ListenAndServe()
}
