package config

import (
	"flag"
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type DBConfig struct {
	DbURL      string `yaml:"database_url" env:"DATABASE_URL"`
	DbPort     string `yaml:"port" env:"DBPORT" env-default:"5432"`
	DbHost     string `yaml:"host" env:"DBHOST" env-default:"localhost"`
	DbName     string `yaml:"name" env:"DBNAME" env-default:"library"`
	DbUser     string `yaml:"user" env:"DBUSER" env-default:"postgres"`
	DbPassword string `yaml:"password" env:"DBPASSWORD"`
}

type AppConfig struct {
	AppPort       string `yaml:"port" env:"PORT"`
	ReqTimeoutSec int    `yaml:"timeout" env:"REQTIMEOUTSEC" env-default:"10"`
}

type LogConfig struct {
	LogLevel string `yaml:"log_level" env:"LOG_LEVEL" env-default:"INFO"`
}

type Config struct {
	DbConf  DBConfig  `yaml:"db"`
	AppConf AppConfig `yaml:"app"`
	LogConf LogConfig `yaml:"logging"`
}

func PrepareConfig() *Config {
	var cfg Config
	configFile := getConfigFile()

	if err := cleanenv.ReadConfig(configFile, &cfg); err != nil {
		fmt.Printf("Unable to get app configuration due to: %s\n", err.Error())
	}
	return &cfg
}

func getConfigFile() string {
	configFile := flag.String("config", "config.yaml", "config file")
	flag.Parse()
	return *configFile
}
