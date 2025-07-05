package main

import (
	"io"

	api "github.com/k-morozov/social/api/restapi"
	"github.com/k-morozov/social/internal/log"
	mware "github.com/k-morozov/social/internal/middleware"
	"github.com/k-morozov/social/internal/service"
	"github.com/labstack/echo/v4"
)

func main() {
	logger := log.NewLogger("debug")

	e := echo.New()

	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())
	// e.Use(middleware.CORS())

	e.Logger.SetOutput(io.Discard)
	e.Use(mware.ZerologMiddleware(logger))

	server := &service.Social{}

	api.RegisterHandlers(e, server)

	e.Logger.Fatal(e.Start(":8080"))
}
