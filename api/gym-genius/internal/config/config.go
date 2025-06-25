package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"os"
)

type Config struct {
	Env  string `yaml:"env"`
	Http Http   `yaml:"http"`
	DB   DB     `yaml:"db"`
}

type Http struct {
	Port    string `yaml:"port"`
	Address string `yaml:"address"`
}

type DB struct {
	Url string `yaml:"url"`
}

func getConfigPath() (string, error) {
	root, err := os.Getwd()
	if err != nil {
		return "", err
	}
	path := root + "/config/config.yaml"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return "", err
	}
	return path, nil
}

func MustLoadConfig() *Config {
	var cfg Config
	path, err := getConfigPath()

	if err != nil {
		panic("config file not found: " + err.Error())
	}

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("failed to read config: " + err.Error())
	}

	return &cfg
}
