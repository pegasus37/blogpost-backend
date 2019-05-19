package config_test

import (
	"os"
	"strconv"
	"testing"

	"github.com/pegasus37/blogpost-backend/config"
	"github.com/stretchr/testify/assert"
)

// func TestSetUpEnv(t *testing.T) {
// 	t.Run("success set env", func(t *testing.T) {
// assert.Equal(t, viper.GetString("PEGASUS_37_DB_HOST"), "35.240.204.148")
// assert.Equal(t, viper.GetInt("PEGASUS_37_DB_PORT"), 5432)
// assert.Equal(t, viper.GetString("PEGASUS_37_DB_USERNAME"), "pegasus37")
// assert.Equal(t, viper.GetString("PEGASUS_37_DB_PASSWORD"), "PegasuSt33ga7")
// assert.Equal(t, viper.GetString("PEGASUS_37_DB_NAME"), "pegasus37")
// assert.Equal(t, viper.GetString("PEGASUS_37_DB_DRIVER"), "postgres")
// 	})
// }

func TestGetConfig(t *testing.T) {
	t.Run("success get config", func(t *testing.T) {
		db_port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
		expected := &config.Configuration{
			DB: &config.DBConfig{
				Host:         os.Getenv("DB_HOST"),
				Port:         db_port,
				Username:     os.Getenv("DB_USERNAME"),
				Password:     os.Getenv("DB_PASSWORD"),
				DatabaseName: os.Getenv("DB_NAME"),
				Driver:       os.Getenv("DB_DRIVER"),
			},
		}

		conf := config.GetConfig()
		assert.Equal(t, conf, expected)
	})
}
