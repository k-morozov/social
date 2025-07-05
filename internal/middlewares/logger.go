package service

import (
	"bytes"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/rs/zerolog"
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

func ZerologMiddleware(logger zerolog.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			logger.Info().
				Str("method", c.Request().Method).
				Str("URI", c.Request().RequestURI).
				Str("remote addr", c.Request().RemoteAddr).
				Msg("request")

			buffer := new(bytes.Buffer)
			writer := io.MultiWriter(c.Response().Writer, buffer)

			c.Response().Writer = &doubleWriter{
				Writer:         writer,
				ResponseWriter: c.Response().Writer,
			}

			if err := next(c); err != nil {
				c.Error(err)
			}

			logger.Info().
				Int("status", c.Response().Status).
				Str("body", buffer.String()).
				Msg("response")
			return nil
		}
	}
}
