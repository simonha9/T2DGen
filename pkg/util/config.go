package util

import (
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yaml"
)

var (
	YelpAPIKey            string
	DefaultYelpAPIKeyPath string
)

type Config struct {
	config *config.Config
}

func init() {
	config.WithOptions(config.ParseEnv)
	config.AddDriver(yaml.Driver)

	err := config.LoadFiles("./config.yaml")
	if err != nil {
		panic(err)
	}
}

func LoadConfig(path string) {
	config.WithOptions(config.ParseEnv)
	config.AddDriver(yaml.Driver)

	err := config.LoadFiles(path)
	if err != nil {
		panic(err)
	}
}

func Get(key string) {

}

func initTest(path string) {
	config.WithOptions(config.ParseEnv)
	config.AddDriver(yaml.Driver)

	err := config.LoadFiles()
	if err != nil {
		panic(err)
	}
}
