package config

import (
	"github.com/spf13/viper"
)

type Configuration struct {
	HTTP_PORT string
	DB        *DBConfig
}

type DBConfig struct {
	Host         string
	Port         int
	Username     string
	Password     string
	DatabaseName string
	Driver       string
}

func setUpEnv() {
	viper.SetEnvPrefix("PEGASUS_37")
	viper.AutomaticEnv()

	viper.BindEnv("PEGASUS_37_API_PORT", "SERVER_PORT")

	viper.BindEnv("PEGASUS_37_DB_HOST", "DB_HOST")
	viper.BindEnv("PEGASUS_37_DB_PORT", "DB_PORT")
	viper.BindEnv("PEGASUS_37_DB_USERNAME", "DB_USERNAME")
	viper.BindEnv("PEGASUS_37_DB_PASSWORD", "DB_PASSWORD")
	viper.BindEnv("PEGASUS_37_DB_NAME", "DB_NAME")
	viper.BindEnv("PEGASUS_37_DB_DRIVER", "DB_DRIVER")
}

func GetConfig() *Configuration {

	setUpEnv()

	DBConfig := &DBConfig{
		Host:         viper.GetString("PEGASUS_37_DB_HOST"),
		Port:         viper.GetInt("PEGASUS_37_DB_PORT"),
		Username:     viper.GetString("PEGASUS_37_DB_USERNAME"),
		Password:     viper.GetString("PEGASUS_37_DB_PASSWORD"),
		DatabaseName: viper.GetString("PEGASUS_37_DB_NAME"),
		Driver:       viper.GetString("PEGASUS_37_DB_DRIVER"),
	}

	return &Configuration{
		HTTP_PORT: viper.GetString("PEGASUS_37_API_PORT"),
		DB:        DBConfig,
	}
}
