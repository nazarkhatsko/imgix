package config

import (
	"errors"
	"os"
	"strconv"
)

const (
	APP_ENV_PROD  = "prod"
	APP_ENV_DEV   = "dev"
	APP_ENV_DEBUG = "debug"
)

type AppConfig struct {
	Name    string `json:"APP_NAME"`    // require param
	Version string `json:"APP_VERSION"` // require param
	Port    int    `json:"APP_PORT"`    // require param
	Env     string `json:"APP_ENV"`     // require param
}

func LoadAppConfig() (AppConfig, error) {
	name := os.Getenv("APP_NAME")
	if name == "" {
		return AppConfig{}, errors.New("APP_NAME is not defined")
	}

	version := os.Getenv("APP_VERSION")
	if version == "" {
		return AppConfig{}, errors.New("APP_VERSION is not defined")
	}

	port, err := strconv.Atoi(os.Getenv("APP_PORT"))
	if err != nil {
		return AppConfig{}, err
	}

	env := os.Getenv("APP_ENV")
	if env == "" {
		return AppConfig{}, errors.New("APP_ENV is not defined")
	}

	conf := AppConfig{
		Name:    name,
		Version: version,
		Port:    port,
		Env:     env,
	}

	return conf, nil
}

func MustLoadAppConfig() AppConfig {
	conf, err := LoadAppConfig()
	if err != nil {
		panic(err)
	}
	return conf
}
