package handlers

import (
	"os"

	"github.com/kzw200015/picbed/cache"
	"github.com/labstack/echo/v4"
)

func HandleRemove(c echo.Context) error {
	id := c.Param("id")

	dstPath, err := cache.Get(id)
	if err != nil {
		return err
	}

	if err = os.Remove(dstPath); err != nil {
		return err
	}

	return c.NoContent(204)
}
