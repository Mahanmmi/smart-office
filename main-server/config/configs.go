package config

import (
	"github.com/spf13/viper"
	"log"
)

type MainServerConfig struct {
	PostgresDB struct {
		Host     string
		Port     uint16
		User     string
		Database string
	}
}

func NewMainServerConfig() *MainServerConfig {
	viper.SetConfigName("config.yaml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("failed to load config file with error: %v", err)
	}

	conf := MainServerConfig{
		PostgresDB: struct {
			Host     string
			Port     uint16
			User     string
			Database string
		}{},
	}
	conf.PostgresDB.Host = viper.GetString("postgres.host")
	conf.PostgresDB.Port = uint16(viper.GetInt("postgres.port"))
	conf.PostgresDB.User = viper.GetString("postgres.user")
	conf.PostgresDB.Database = viper.GetString("postgres.database")

	return &conf
}
