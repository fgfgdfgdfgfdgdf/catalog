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
}

type Postgres struct {
	USER                    string
	PASSWORD                string
	HOST                    string `yaml:"host"`
	NAME                    string
	PORT                    string        `yaml:"port"`
	CONTEXT_CANCEL_SECONDS  int64         `yaml:"context_cancel_seconds"`
	CONTEXT_CANCEL_DURATION time.Duration `yaml:"-"`
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
	}

	err = utils.ReadYaml(c, path)
	if err != nil {
		panic(err)
	}

	c.Postgres.CONTEXT_CANCEL_DURATION = time.Duration(c.Postgres.CONTEXT_CANCEL_SECONDS)

	c.Postgres.PASSWORD = os.Getenv("POSTGRES_PASSWORD")
	c.Postgres.USER = os.Getenv("POSTGRES_USER")
	c.Postgres.NAME = os.Getenv("POSTGRES_DB")

	config = c

	return config.App
}

func Pg() *Postgres {
	return config.Postgres
}

type MainConfig interface {
	HOST() string
	PORT() string
}

func App() *Application {
	return config.App
}
