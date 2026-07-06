package yaml

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Blabla string `yaml:"blabla"`
}

func (config *Config) init() *Config {
	yamlFile, error := os.ReadFile("config.yaml")
	if error != nil {
		panic(error)
	}
	error = yaml.Unmarshal(yamlFile, config)
	if error != nil {
		panic(error)
	}

	return config
}

func InitConfig() *Config {
	var config Config
	config.init()

	return &config
}
