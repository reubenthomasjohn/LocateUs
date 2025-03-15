package util

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	AccountSid string `mapstructure:"TWILIO_ACCOUNT_SID"`
	AuthToken string `mapstructure:"TWILIO_AUTH_TOKEN"`
	SenderNumber string `mapstructure:"SENDER_NUMBER"`
	DBDriver string `mapstructure:"DB_DRIVER"`
	DBSource string `mapstructure:"DB_SOURCE"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
	NgrokUrl string `mapstructure:"NGROK_URL"`
	DomainName string `mapstructure:"DOMAIN_NAME"`
	PrefixUrl	string
	TokenSymmetricKey string `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
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

	if config.DomainName != "" {
		config.PrefixUrl = config.DomainName
	} else if config.NgrokUrl != "" {
		config.PrefixUrl = config.NgrokUrl
	}
	return
}