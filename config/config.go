// Package config for environment configuration details
package config

import (
	"flag"
	"log"

	"github.com/spf13/viper"
)

var (
	environment = flag.String("env", "dev", "deployment environment for config. Default `env`")
	appName     string
	env         *EnvConfig
	postgres    *PostgresConfig
)

func init() {
	flag.Parse()
	loadConfig()
}

func loadConfig() {
	viper.SetConfigName("config-" + *environment)
	viper.AddConfigPath("./")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("err while reading config file", err)
	}
	if err := viper.Sub("env").Unmarshal(&env); err != nil {
		log.Fatal("err while marshalling env config", err)
	}
	if err := viper.Sub("postgres").Unmarshal(&postgres); err != nil {
		log.Fatal("err while marshalling postgres config", err)
	}
	appName = viper.GetString("name")
}

// Env returns environment config defined in env specific config file
func Env() EnvConfig {
	return *env
}

// AppName returns application name
func AppName() string {
	return appName
}

// Postgres returns postgres config defined in env specific config file
func Postgres() *PostgresConfig {
	return postgres
}
