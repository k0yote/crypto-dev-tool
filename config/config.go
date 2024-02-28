package config

import (
	"github.com/k0yote/backend-wallet/util"
	"github.com/spf13/viper"
)

// Config stores all configuration of the application.
// The values are read by viper from a config file or environment variable.
type Config struct {
	Environment          string `mapstructure:"ENVIRONMENT"`
	HTTPServerAddress    string `mapstructure:"HTTP_SERVER_ADDRESS"`
	OTPIssuer            string `mapstructure:"OTP_ISSURE"`
	PassCodeExpirePeriod int    `mapstructure:"PASSCODE_EXPIRE_PERIOD"`
	EmailSenderAddress   string `mapstructure:"EMAIL_SENDER_ADDRESS"`
	EmailSenderPassword  string `mapstructure:"EMAIL_SENDER_PASSWORD"`
	SmtpAuthAddress      string `mapstructure:"SMTP_AUTH_ADDRESS"`
	SmtpServerAddress    string `mapstructure:"SMTP_SERVER_ADDRESS"`
	GoogleCredentialPath string `mapstructure:"GOOGLE_CREDENTIAL_PATH"`
	GoogleKmsLocation    string `mapstructure:"GOOGLE_KMS_LOCATION"`
}

func LoadConfig() (Config, error) {
	var config Config
	path, err := util.GetRootPath()
	if err != nil {
		return config, err
	}

	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return config, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}

	return config, nil
}
