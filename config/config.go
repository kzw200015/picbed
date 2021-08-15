package config

import (
	"flag"
	"os"

	"gopkg.in/yaml.v3"
)

const (
	IMG_ROUTE_PREFIX = "/images"
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

func initDefault() {
	config.Port = 8080
	config.ImgSrc = "./images"
	config.Debug = false
	config.Referers = []string{"*"}
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
	initDefault()

	path := flag.String("c", "", "config file path")

	flag.Parse()

	if !(*path == "") {
		initFromFile(*path)
	}
}
