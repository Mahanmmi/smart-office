package config

import (
	"github.com/spf13/viper"
	"log"
)

type LocalServerConfig struct {
	PostgresDB struct {
		Host     string
		Port     uint16
		User     string
		Password string
		Database string
	}
	MQTT struct{
		Broker 	 string
		Port 	 uint16
		ClientID string
	}
	// HTTPServerPort string
	// JWTSecret      []byte
	OfficeAPIKey string
}

func NewLocalServerConfig() *LocalServerConfig {
	viper.SetConfigName("config.yaml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("..")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("failed to load config file with error: %v", err)
	}

	conf := LocalServerConfig{
		PostgresDB: struct {
			Host     string
			Port     uint16
			User     string
			Password string
			Database string
		}{},
		MQTT: struct {
			Broker 	 string
			Port 	 uint16
			ClientID string
		}{},
	}
	conf.PostgresDB.Host = viper.GetString("postgres.host")
	conf.PostgresDB.Port = uint16(viper.GetInt("postgres.port"))
	conf.PostgresDB.User = viper.GetString("postgres.user")
	conf.PostgresDB.Password = viper.GetString("postgres.password")
	conf.PostgresDB.Database = viper.GetString("postgres.database")


	conf.MQTT.Broker = viper.GetString("mqtt.broker")
	conf.MQTT.Port = uint16(viper.GetInt("mqtt.port"))
	conf.MQTT.ClientID = viper.GetString("mqtt.clientid")

	conf.OfficeAPIKey = viper.GetString("office.apikey")

	// conf.HTTPServerPort = viper.GetString("http.port")
	// conf.JWTSecret = []byte(viper.GetString("http.jwt_secret"))

	// conf.OfficeKeyIDMap = viper.GetStringMapString("offices")

	return &conf
}
