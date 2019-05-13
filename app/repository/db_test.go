package repository_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pegasus37/blogpost-backend/app/repository"
	"github.com/pegasus37/blogpost-backend/config"
)

func TestSetUpDBConnection(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		config := &config.Configuration{
			DB: &config.DBConfig{
				Host:         "35.240.204.148",
				Port:         5432,
				Username:     "pegasus37",
				Password:     "PegasuSt33ga7",
				DatabaseName: "pegasus37",
				Driver:       "postgres",
			},
		}

		DbHandler := repository.SetUpDBConnection(config)
		assert.NotNil(t, DbHandler)
	})

	t.Run("negative", func(t *testing.T) {
		config := &config.Configuration{
			DB: &config.DBConfig{
				Host:         "",
				Port:         0,
				Username:     "",
				Password:     "",
				DatabaseName: "",
				Driver:       "mysql",
			},
		}

		DbHandler := repository.SetUpDBConnection(config)
		assert.Nil(t, DbHandler)
	})
}
