package config

import (
	"errors"
	"os"
	"strconv"
)

type DatabaseConfig struct {
	Host     string `json:"DATABASE_HOST"`     // require param
	Port     int    `json:"DATABASE_PORT"`     // require param
	Name     string `json:"DATABASE_NAME"`     // require param
	User     string `json:"DATABASE_USER"`     // require param
	Password string `json:"DATABASE_PASSWORD"` // require param
}

func LoadDatabaseConfig() (DatabaseConfig, error) {
	host := os.Getenv("DATABASE_HOST")
	if host == "" {
		return DatabaseConfig{}, errors.New("DATABASE_HOST is not defined")
	}

	port, err := strconv.Atoi(os.Getenv("DATABASE_PORT"))
	if err != nil {
		return DatabaseConfig{}, err
	}

	user := os.Getenv("DATABASE_USER")
	if user == "" {
		return DatabaseConfig{}, errors.New("DATABASE_USER is not defined")
	}

	password := os.Getenv("DATABASE_PASSWORD")
	if password == "" {
		return DatabaseConfig{}, errors.New("DATABASE_PASSWORD is not defined")
	}

	name := os.Getenv("DATABASE_NAME")
	if name == "" {
		return DatabaseConfig{}, errors.New("DATABASE_NAME is not defined")
	}

	conf := DatabaseConfig{
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
		Name:     name,
	}

	return conf, nil
}

func MustLoadDatabaseConfig() DatabaseConfig {
	conf, err := LoadDatabaseConfig()
	if err != nil {
		panic(err)
	}
	return conf
}
