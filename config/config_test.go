package config_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pegasus37/blogpost-backend/config"
)

func TestSetUpConfig(t *testing.T) {
	t.Run("success get config", func(t *testing.T) {
		expected := &config.Configuration{
			DB: &config.DBConfig{
				Host:         "35.240.204.148",
				Port:         5432,
				Username:     "pegasus37",
				Password:     "PegasuSt33ga7",
				DatabaseName: "pegasus37",
				Driver:       "postgres",
			},
		}

		conf := config.GetConfig()
		assert.Equal(t, conf, expected)
	})
}
