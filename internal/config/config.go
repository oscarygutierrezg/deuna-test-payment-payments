package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	BootstrapServers string
	GroupID          string
	ConsumerTopic    string
	ProducerTopic    string
	DBUser           string
	DBPassword       string
	DBName           string
	DBHost           string
	DBPort           int64
}

func LoadConfig() (*Config, error) {
	viper.SetDefault("bootstrapServers", "localhost:9092")
	viper.SetDefault("groupId", "payment-payments-group")
	viper.SetDefault("consumerTopic", "com.deuna.payment.payment.v1.payments.updated")
	viper.SetDefault("producerTopic", "com.deuna.payment.payment_banking_x.v1.payments")

	viper.SetDefault("dbUser", "root")
	viper.SetDefault("dbPassword", "root")
	viper.SetDefault("dbName", "payments")
	viper.SetDefault("dbHost", "localhost")
	viper.SetDefault("dbPort", int64(5432))

	viper.AutomaticEnv()

	config := &Config{
		BootstrapServers: viper.GetString("bootstrapServers"),
		GroupID:          viper.GetString("groupId"),
		ConsumerTopic:    viper.GetString("consumerTopic"),
		ProducerTopic:    viper.GetString("producerTopic"),

		DBUser:     viper.GetString("dbUser"),
		DBPassword: viper.GetString("dbPassword"),
		DBName:     viper.GetString("dbName"),
		DBHost:     viper.GetString("dbHost"),
		DBPort:     viper.GetInt64("dbPort"),
	}

	return config, nil
}
