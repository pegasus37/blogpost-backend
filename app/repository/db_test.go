package repository_test

import (
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pegasus37/blogpost-backend/app/repository"
	"github.com/pegasus37/blogpost-backend/config"
)

func TestSetUpDBConnection(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		db_port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
		config := &config.Configuration{
			DB: &config.DBConfig{
				Host:         os.Getenv("DB_HOST"),
				Port:         db_port,
				Username:     os.Getenv("DB_USERNAME"),
				Password:     os.Getenv("DB_PASSWORD"),
				DatabaseName: os.Getenv("DB_NAME"),
				Driver:       os.Getenv("DB_DRIVER"),
			},
		}

		DbHandler := repository.SetUpDBConnection(config)
		assert.NotNil(t, DbHandler)
	})
}
