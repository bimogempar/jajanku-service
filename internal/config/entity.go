package config

import (
	"errors"
	"os"

	dotenv "github.com/golobby/dotenv"
)

type LoadConfig struct {
	Database struct {
		Driver   string `env:"DB_DRIVER"`
		Host     string `env:"DB_HOST"`
		Name     string `env:"DB_NAME"`
		Username string `env:"DB_USER"`
		Password string `env:"DB_PASSWORD"`
		Port     string `env:"DB_PORT"`
	}
	JWTConfig struct {
		SecretKey string `env:"JWT_SECRET_KEY`
	}
}

var Config LoadConfig

func New() (LoadConfig, error) {
	//Load Environment file
	file, err := os.Open(".env")
	if err != nil {
		return Config, errors.New("error loading .env file")
	}

	err = dotenv.NewDecoder(file).Decode(&Config)
	if err != nil {
		return Config, errors.New("cannot decode .env file")
	}
	return Config, err
}
