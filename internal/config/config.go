package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"server"`
	DB struct {
		Path      string `yaml:"path"`
		PathToSQL string `yaml:"path_to_sql"`
	} `yaml:"db"`
	JWT struct {
		PathToKey string `yaml:"path_to_key"`
		Access    struct {
			Time int `yaml:"time"`
		} `yaml:"access"`
		Refresh struct {
			Time int `yaml:"time"`
		} `yaml:"refresh"`
	} `yaml:"jwt"`
}

func LoadConfig(path string) (*Config, error) {

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var cfg Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
