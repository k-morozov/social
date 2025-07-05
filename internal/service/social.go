package service

import (
	openapi "github.com/k-morozov/social/api/restapi"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
)

type social struct {
	logger zerolog.Logger
}

var _ openapi.ServerInterface = &social{}

func NewSocial(logger zerolog.Logger) openapi.ServerInterface {
	return &social{
		logger: logger,
	}
}

func (s *social) PostLogin(ctx echo.Context) error              { return nil }
func (s *social) PostUserRegister(ctx echo.Context) error       { return nil }
func (s *social) GetUserGetId(ctx echo.Context, id int64) error { return nil }
