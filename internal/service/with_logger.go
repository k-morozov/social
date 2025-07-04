package service

import (
	"bytes"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/k-morozov/social/internal/log"
)

type doubleWriter struct {
	io.Writer
	http.ResponseWriter
}

func (w *doubleWriter) WriteHeader(code int) {
	w.ResponseWriter.WriteHeader(code)
}

func (w *doubleWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

func OptionLogger() ServiceHTTPOption {
	return func(s *ServiceHTTP) {
		withLogger(s)
	}
}

func withLogger(s *ServiceHTTP) {
	s.logger = log.NewLogger("debug")

	s.engine.Use(logResponse(s))
	s.engine.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogMethod: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			s.logger.Info().
				Str("method", v.Method).
				Str("URI", v.URI).
				Int("status", v.Status).
				Msg("request")

			return nil
		},
	}))
}

func logResponse(s *ServiceHTTP) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		s.logger.Debug().Msg("Call LogResponse")
		return func(c echo.Context) error {

			buffer := new(bytes.Buffer)
			writer := io.MultiWriter(c.Response().Writer, buffer)

			c.Response().Writer = &doubleWriter{
				Writer:         writer,
				ResponseWriter: c.Response().Writer,
			}

			if err := next(c); err != nil {
				c.Error(err)
			}

			s.logger.Info().
				Int("status", c.Response().Status).
				Str("body", buffer.String()).
				Msg("response")
			return nil
		}
	}
}
