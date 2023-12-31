package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type App struct {
	Port string `json:"port" yaml:"port"`
}

type Config struct {
	App App `json:"app" yaml:"app"`
}

func ReadConfig(filename string) (*Config, error) {
	conf := &Config{}
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	if yaml.Unmarshal(content, conf) != nil {
		panic(fmt.Sprintf("%s: %v", filename, err))
	}

	return conf, nil
}
