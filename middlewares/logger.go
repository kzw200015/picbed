package middlewares

import (
	"time"

	"github.com/labstack/echo/v4"
)

func Logger() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			isErr := false

			start := time.Now()
			err := next(c)
			if err != nil {
				c.Error(err)
				isErr = true
			}
			end := time.Now()

			latency := end.Sub(start).String()
			status := c.Response().Status
			method := c.Request().Method
			uri := c.Request().RequestURI

			if isErr {
				switch {
				case status >= 500:
					c.Logger().Error(err)
				case status >= 400:
					c.Logger().Warn(err)
				}
			}

			c.Logger().Infof("%s %d %s %s", latency, status, method, uri)

			return err
		}
	}
}
