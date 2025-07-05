package service

import (
	api "github.com/k-morozov/social/api/restapi"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
)

type Social struct {
	logger zerolog.Logger
}

var _ api.ServerInterface = &Social{}

func (s *Social) PostLogin(ctx echo.Context) error              { return nil }
func (s *Social) PostUserRegister(ctx echo.Context) error       { return nil }
func (s *Social) GetUserGetId(ctx echo.Context, id int64) error { return nil }
