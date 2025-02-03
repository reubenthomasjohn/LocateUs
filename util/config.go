package util

import "github.com/spf13/viper"

type Config struct {
	AccountSid string `mapstructure:"TWILIO_ACCOUNT_SID"`
	AuthToken string `mapstructure:"TWILIO_AUTH_TOKEN"`
	SenderNumber string `mapstructure:"SENDER_NUMBER"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return 
	}

	err = viper.Unmarshal(&config)
	return
}