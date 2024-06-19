package config

import (
	"os"

	"github.com/joho/godotenv"
)

type (
	Config struct {
		ServerHost string
		OwnHost    string
	}
)

func (c *Config) Load() error {
	if err := godotenv.Load(); err != nil {
		return err
	}
	c.ServerHost = os.Getenv("SERVER_HOST")
	c.OwnHost = os.Getenv("OWN_HOST")
	return nil
}

func New() (*Config, error) {
	var cnfg Config
	if err := cnfg.Load(); err != nil {
		return nil, err
	}
	return &cnfg, nil
}
