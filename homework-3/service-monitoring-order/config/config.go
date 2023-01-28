package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"strings"
)

type Config struct {
	App   `yaml:"app"`
	PG    `yaml:"postgres"`
	Kafka `yaml:"kafka"`
}

type App struct {
	Name    string `yaml:"name" env:"APP-NAME"`
	Host    string `yaml:"host" env:"APP-HOST"`
	Version string `yaml:"version" env:"APP-VERSION"`
}

type PG struct {
	Pool     int    `yaml:"pool" env:"PG-POOL"`
	Host     string `yaml:"host" env:"PG-HOST"`
	Port     string `yaml:"port" env:"PG-PORT"`
	Dbname   string `yaml:"db_name" env:"PG-DBNAME"`
	Username string `yaml:"username" env:"PG-USER"`
	Password string `yaml:"password" env:"PG-PASSWORD"`
}

type Kafka struct {
	Brokers          []string `yaml:"brokers" env:"KAFKA-BROKERS"`
	GroupID          string   `yaml:"group_id" env:"KAFKA-GROUP"`
	IncomeTopic      string   `yaml:"income_topic" env:"KAFKA-INCOME-TOPIC"`
	ResetTopic       string   `yaml:"reset_topic" env:"KAFKA-RESET-TOPIC"`
	ReservationTopic string   `yaml:"reservation_topic" env:"KAFKA-RESERVATION-TOPIC"`
}

func New(path string) (*Config, error) {
	var cfg Config

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		return nil, err
	}

	// override brokers
	// env variables cannot be slices, only strings
	if len(cfg.Brokers) == 1 {
		cfg.Brokers = strings.Split(cfg.Brokers[0], ",")
	}

	return &cfg, nil
}
