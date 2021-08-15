package middlewares

import (
	"net/url"
	"path/filepath"

	"github.com/kzw200015/picbed/config"
	"github.com/labstack/echo/v4"
)

func Referer(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		for _, referer := range config.Get().Referers {
			u, err := url.Parse(c.Request().Referer())
			if err != nil {
				return err
			}

			ok, err := filepath.Match(referer, u.Host)
			if err != nil {
				return err
			}

			if ok {
				return next(c)
			}
		}

		return echo.ErrForbidden
	}
}
