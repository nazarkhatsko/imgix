package config

import (
	"flag"

	"github.com/joho/godotenv"
)

func LoadConfig(envFilenames ...string) error {
	configPath := flag.String("config", ".env", "path to the config file")
	flag.Parse()

	envFilenames = append(envFilenames, *configPath)
	err := godotenv.Load(envFilenames...)
	if err != nil {
		return err
	}
	return nil
}

func MustLoadConfig(envFilenames ...string) {
	err := LoadConfig(envFilenames...)
	if err != nil {
		panic(err)
	}
}

func InitConfig() {}
