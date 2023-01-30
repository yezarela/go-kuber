package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	AppPort    string `env:"APP_PORT" env-default:"4001"`
	DBHost     string `env:"DB_HOST" env-required:"true"`
	DBPort     string `env:"DB_PORT" env-required:"true"`
	DBUser     string `env:"DB_USER" env-required:"true"`
	DBPassword string `env:"DB_PASSWORD" env-required:"true"`
	DBName     string `env:"DB_NAME" env-required:"true"`
}

var Data Config

func Load(path string) error {
	err := cleanenv.ReadConfig(path, &Data)
	if err != nil {
		return err
	}

	return nil
}
