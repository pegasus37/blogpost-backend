package repository

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	_ "github.com/lib/pq"

	"github.com/pegasus37/blogpost-backend/config"
	"github.com/pegasus37/blogpost-backend/entity"
)

type MySqlRepo struct {
	Db *sql.DB
}

func NewMySqlRepo(config *config.DBConfig) *MySqlRepo {
	user := config.Username
	pass := config.Password
	host := config.Host
	port := config.Port
	name := config.DatabaseName
	address := host + ":" + strconv.Itoa(port)
	dsn := fmt.Sprintf(config.Driver+"://%[1]s:%[2]s@%[3]s/%[4]s?sslmode=%[5]s", user, pass, address, name, DefaultSSLMode)

	db, err := sql.Open(config.Driver, dsn)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
		return nil
	}

	db.SetMaxOpenConns(MaxOpenConnection)

	mySqlRepo := &MySqlRepo{
		Db: db,
	}

	return mySqlRepo
}

func (p MySqlRepo) Get() *entity.Pegasus {
	pegasus := &entity.Pegasus{
		Name:   "Pegasus 36",
		Number: 36,
		Type:   "Air Zoom",
		Year:   2019,
	}

	return pegasus
}
