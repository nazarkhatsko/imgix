package config

import (
	"errors"
	"os"
	"strconv"
)

type RedisConfig struct {
	Host     string `json:"REDIS_HOST"`    // require param
	Port     int    `json:"REDIS_PORT"`    // require param
	Password string `json:"REDIS_PASWORD"` // optional param
}

func LoadRedisConfig() (RedisConfig, error) {
	host := os.Getenv("REDIS_HOST")
	if host == "" {
		return RedisConfig{}, errors.New("REDIS_HOST is not defined")
	}

	port, err := strconv.Atoi(os.Getenv("REDIS_PORT"))
	if err != nil {
		return RedisConfig{}, err
	}

	password := os.Getenv("REDIS_PASWORD")

	conf := RedisConfig{
		Host:     host,
		Port:     port,
		Password: password,
	}

	return conf, nil
}

func MustLoadRedisConfig() RedisConfig {
	conf, err := LoadRedisConfig()
	if err != nil {
		panic(err)
	}
	return conf
}
