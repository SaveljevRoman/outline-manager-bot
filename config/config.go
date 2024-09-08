package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"outline-manager-bot/pkg/fileloader"
)

const logPoint = "point: init config"

type Config struct {
	TgConf *TGConfig
	PgConf *PostgresConfig
}

type TGConfig struct {
	TgBotToken    string `env:"TG_BOT_TOKEN" env-default:""`
	Debug         bool   `env:"TG_BOT_DEBUG" env-default:"false"`
	UpdateOffset  int    `env:"TG_BOT_UPDATE_OFFSET" env-default:"0"`
	UpdateTimeout int    `env:"TG_BOT_UPDATE_TIMEOUT" env-default:"20"`
}

type PostgresConfig struct {
	Host    string `env:"POSTGRES_HOST" env-default:"127.0.0.1"`
	Port    string `env:"POSTGRES_PORT" env-default:"54321"`
	User    string `env:"POSTGRES_USERNAME" env-default:"postgres"`
	Pass    string `env:"POSTGRES_PASSWORD" env-default:"root"`
	DbName  string `env:"POSTGRES_DATABASE" env-default:"outline_manager"`
	SslMode string `env:"POSTGRES_SSL_MODE" env-default:"disable"`
}

func LoadConfig() *Config {
	fileLoader.EnvLoader()

	var tg TGConfig
	if err := cleanenv.ReadEnv(&tg); err != nil {
		fmt.Println(logPoint)
		log.Fatal(err)
	}

	var pg PostgresConfig
	if err := cleanenv.ReadEnv(&pg); err != nil {
		fmt.Println(logPoint)
		log.Fatal(err)
	}

	return &Config{
		TgConf: &tg,
		PgConf: &pg,
	}
}
