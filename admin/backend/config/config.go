package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)


type Config struct {
	Secret     string `envconfig:"AUTH_SECRET" required:"true"`
	JwtSecret  string `envconfig:"JWT_SECRET" required:"true"`
	DbHost     string `envconfig:"DB_HOST" required:"true"`
	DbUser     string `envconfig:"DB_USER" required:"true"`
	DbPassword string `envconfig:"DB_PASSWORD" required:"true"`
	DbName     string `envconfig:"DB_NAME" required:"true"`
	DbPort     string `envconfig:"DB_PORT" required:"true"`
}

var C Config

func init() {
	// does not override set env variables
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	if err := envconfig.Process("", &C); err != nil {
		panic(err)
	}
}
