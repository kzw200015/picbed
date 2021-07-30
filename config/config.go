package config

import (
	"flag"
	"os"
	"strconv"
	"strings"

	"gopkg.in/yaml.v3"
)

const (
	IMG_ROUTE_PREFIX = "/images"
	ENV_NAME_PORT    = "PB_PORT"
	ENV_NAME_IMG_SRC = "PB_PORT_IMG_SRC"
	ENV_NAME_DEBUG   = "PB_PORT_DEBUG"
)

var config Config

type Config struct {
	Port     int      `json:"port,omitempty" yaml:"port,omitempty"`
	ImgSrc   string   `json:"imgSrc,omitempty" yaml:"imgSrc,omitempty"`
	Debug    bool     `json:"debug,omitempty" yaml:"debug,omitempty"`
	Referers []string `json:"referers,omitempty" yaml:"referers,omitempty"`
}

func Get() Config {
	return config
}

func initFromFile(path string) {
	bs, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(bs, &config)
	if err != nil {
		panic(err)
	}
}

func init() {

	port := flag.Int("p", 8080, "port")
	imgSrc := flag.String("s", "./images", "image source")
	debug := flag.Bool("d", false, "debug mode")
	path := flag.String("c", "", "config file path")

	flag.Parse()

	config.Port = *port
	config.ImgSrc = *imgSrc
	config.Debug = *debug
	config.Referers = []string{"*"}

	if !(*path == "") {
		initFromFile(*path)
	}

	if e, err := strconv.Atoi(os.Getenv(ENV_NAME_PORT)); err == nil {
		config.Port = e
	}

	if e := os.Getenv(ENV_NAME_IMG_SRC); e != "" {
		config.ImgSrc = e
	}

	if strings.ToLower(os.Getenv(ENV_NAME_DEBUG)) == "true" {
		config.Debug = true
	}
}
