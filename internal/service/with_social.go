package service

import (
	"github.com/labstack/echo/v4"

	"github.com/k-morozov/social/internal/api"
)

func OptionSocial() ServiceHTTPOption {
	return func(s *ServiceHTTP) {
		configureHandlers(s)
	}
}

func configureHandlers(s *ServiceHTTP) {
	s.api = api.NewSocailAPI(&s.logger)

	s.engine.POST("/user/register", func(c echo.Context) error {
		s.logger.Debug().Msg("Call register")
		s.api.Register(c.Response().Writer, c.Request())
		return nil
	})
}
