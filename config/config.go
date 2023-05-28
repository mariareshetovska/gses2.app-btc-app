package config

import (
	"os"

	"github.com/joho/godotenv"
)

type ConfigEmail struct {
	Host     string
	Port     string
	Username string
	Password string
}

func LoadEnv() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}
	return nil
}

func LoadEmailConfig() (ConfigEmail, error) {
	err := LoadEnv()
	if err != nil {
		return ConfigEmail{}, err
	}

	cnf := ConfigEmail{
		Host:     os.Getenv("EMAIL_HOST"),
		Port:     os.Getenv("EMAIL_PORT"),
		Username: os.Getenv("EMAIL_ADDRESS"),
		Password: os.Getenv("EMAIL_PASSWORD"),
	}

	return cnf, nil
}
