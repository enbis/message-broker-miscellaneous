package config

import (
	"log"

	"github.com/spf13/viper"
)

func InitConfig() error {
	viper.SetDefault("nats_url", "nats://localhost:4222")
	viper.SetDefault("grpc_port", 7777)
	viper.SetDefault("topic", "topic")
	viper.SetDefault("payload", "payload")

	viper.AddConfigPath("./")
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		log.Printf("Using config file: %s", viper.ConfigFileUsed())
	}

	return nil
}
