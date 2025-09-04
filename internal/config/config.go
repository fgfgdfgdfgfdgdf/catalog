package config

import (
	"fmt"
	"os"
	"time"

	"github.com/fgfgdfgdfgfdgdf/catalog/internal/utils"
	"github.com/joho/godotenv"
)

const (
	path = "configs/main.yaml"
)

var (
	config = &configuration{}
)

type configuration struct {
	Postgres *Postgres    `yaml:"postgres"`
	App      *Application `yaml:"app"`
	Redis    *Redis       `yaml:"redis"`
}

type Postgres struct {
	USER                  string
	PASSWORD              string
	HOST                  string `yaml:"host"`
	NAME                  string
	PORT                  string        `yaml:"port"`
	ContextCancelSeconds  int64         `yaml:"context_cancel_seconds"`
	ContextCancelDuration time.Duration `yaml:"-"`
}

type Redis struct {
	USER                  string
	PASSWORD              string
	HOST                  string        `yaml:"host"`
	ContextCancelSeconds  int64         `yaml:"context_cancel_seconds"`
	ContextCancelDuration time.Duration `yaml:"-"`
	Key                   int64         `yaml:"context_cancel_seconds"`
	KeyExpirationDuration time.Duration `yaml:"-"`
	KeyExpirationSeconds  int64         `yaml:"key_expiration_seconds"`
}

type Application struct {
	HOST           string `yaml:"host"`
	PORT           string `yaml:"port"`
	DefaultPerPage int64  `yaml:"default_per_page"`
	DefaultPage    int64  `yaml:"default_page"`
	MaxPerPage     int64  `yaml:"max_per_page"`
}

func Init() *Application {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(".env файл не найден")
	}

	c := &configuration{
		Postgres: &Postgres{},
		Redis:    &Redis{},
	}

	err = utils.ReadYaml(c, path)
	if err != nil {
		panic(err)
	}

	c.Postgres.ContextCancelDuration = time.Duration(c.Postgres.ContextCancelSeconds)
	c.Redis.ContextCancelDuration = time.Duration(c.Redis.ContextCancelSeconds)
	c.Redis.KeyExpirationDuration = time.Duration(c.Redis.KeyExpirationSeconds)

	c.Postgres.PASSWORD = os.Getenv("POSTGRES_PASSWORD")
	c.Postgres.USER = os.Getenv("POSTGRES_USER")
	c.Postgres.NAME = os.Getenv("POSTGRES_DB")

	c.Redis.PASSWORD = os.Getenv("REDIS_USER_PASSWORD")
	c.Redis.USER = os.Getenv("REDIS_USER")

	config = c

	return config.App
}

func Pg() *Postgres {
	return config.Postgres
}

func Rds() *Redis {
	return config.Redis
}

func App() *Application {
	return config.App
}
