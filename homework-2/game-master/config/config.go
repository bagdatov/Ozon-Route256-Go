package config

import (
	"gopkg.in/yaml.v2"
	"os"
	"time"
)

type Config struct {
	App      `yaml:"app"`
	Log      `yaml:"logger"`
	PG       `yaml:"postgres"`
	ChgkBase `yaml:"chgk_base"`
	GRPC     `yaml:"grpc"`
	Gateway  `yaml:"gateway"`
}

type App struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
}

type Log struct {
	Level string `yaml:"log_level"`
}

type Gateway struct {
	Port string `yaml:"port"`
}

type PG struct {
	Pool     int    `yaml:"pool"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Dbname   string `yaml:"db_name"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type GRPC struct {
	Host string `yaml:"host"`
}

type ChgkBase struct {
	Url     string        `yaml:"url"`
	Timeout time.Duration `yaml:"timeout"`
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
