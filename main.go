package main

import (
	"strconv"

	"github.com/kzw200015/picbed/config"
	"github.com/kzw200015/picbed/routes"
)

func main() {
	e := routes.Get()
	e.Debug = config.Get().Debug
	e.HideBanner = true
	e.Logger.SetHeader("[${time_rfc3339}] ${level}")
	e.Logger.Fatal(e.Start(":" + strconv.Itoa(config.Get().Port)))
}
