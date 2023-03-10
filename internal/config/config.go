package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Server Server   `mapstructure:",squash"`
	Kafka  Kafka    `mapstructure:",squash"`
	PSQL   Postgres `mapstructure:",squash"`
}

type Kafka struct {
	Servers    string
	Mechanisms string
	Protocol   string
	Username   string
	Password   string
	CaFile     string
	Group      string
	Topics     string
	Commit     bool
	Offset     string
}

type Server struct {
	Port string
}

type Postgres struct {
	PgUser     string
	PgPassword string
	PgHost     string
	PgPort     string
	PgDb       string
}

func (c *Config) InitCfg() error {
	viper.AddConfigPath("./")
	viper.SetConfigName("config")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	err = viper.Unmarshal(&c)

	return err
}
