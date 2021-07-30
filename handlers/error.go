package handlers

import (
	"github.com/labstack/echo/v4"
	"strconv"
	"strings"
)

func HandleError(err error, c echo.Context) {
	he, ok := err.(*echo.HTTPError)
	if !ok || strings.HasPrefix(strconv.Itoa(he.Code), "5") {
		c.Logger().Error(err.Error())
	} else {
		c.Logger().Warn(err.Error())
	}
	c.Echo().DefaultHTTPErrorHandler(err, c)
}
