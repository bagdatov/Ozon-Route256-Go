package config

import (
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	App      `yaml:"app"`
	Log      `yaml:"logger"`
	GRPC     `yaml:"grpc"`
	Telegram `yaml:"telegram"`
}

type App struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
}

type Log struct {
	Level string `yaml:"log_level"`
}

type GRPC struct {
	Host string `yaml:"host"`
}

type Telegram struct {
	Key string `yaml:"key"`
}

func New(path string) (*Config, error) {
	cfg := &Config{}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(data, cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
