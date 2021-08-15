package main

import (
	"strconv"

	"github.com/kzw200015/picbed/config"
	"github.com/kzw200015/picbed/routes"
	"github.com/labstack/gommon/log"
)

func main() {
	e := routes.Get()
	e.Debug = config.Get().Debug
	e.Logger.SetLevel(log.DEBUG)
	e.HideBanner = true
	e.Logger.SetHeader("[${time_rfc3339}] ${level}")
	e.Logger.Fatal(e.Start(":" + strconv.Itoa(config.Get().Port)))
}
