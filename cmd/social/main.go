package main

import (
	"io"

	"github.com/labstack/echo/v4"

	openapi "github.com/k-morozov/social/api/restapi"
	"github.com/k-morozov/social/internal/log"
	mware "github.com/k-morozov/social/internal/middlewares"
	"github.com/k-morozov/social/internal/service"
)

func main() {
	logger := log.NewLogger("debug")
	e := echo.New()

	// e.Use(middleware.Recover())
	// e.Use(middleware.CORS())

	e.Logger.SetOutput(io.Discard)
	e.Use(mware.ZerologMiddleware(logger))

	server := service.NewSocial(logger)

	openapi.RegisterHandlers(e, server)

	e.Logger.Fatal(e.Start(":8080"))
}
