package config

import (
	"context"
	"os"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/pkg/errors"
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

func LoadFromSecretManager(name string) error {
	ctx := context.Background()
	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to setup secretmanager client")
	}
	defer client.Close()

	accessRequest := &secretmanagerpb.AccessSecretVersionRequest{
		Name: name,
	}

	result, err := client.AccessSecretVersion(ctx, accessRequest)
	if err != nil {
		return errors.Wrap(err, "failed to access secret version")
	}

	err = os.WriteFile(".env", result.Payload.Data, 0644)
	if err != nil {
		return errors.Wrap(err, "failed to write secret to file")
	}

	return Load(".env")
}
