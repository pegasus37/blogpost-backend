package repository

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"github.com/pegasus37/blogpost-backend/config"
)

const (
	DefaultSSLMode    = "require"
	MaxOpenConnection = 1
)

type DBHandler struct {
	Db *sql.DB
}

func SetUpDBConnection(config *config.Configuration) *DBHandler {
	dbConf := config.DB
	user := dbConf.Username
	pass := dbConf.Password
	host := dbConf.Host
	port := dbConf.Port
	name := dbConf.DatabaseName
	address := host + ":" + strconv.Itoa(port)
	dsn := fmt.Sprintf(dbConf.Driver+"://%[1]s:%[2]s@%[3]s/%[4]s?sslmode=%[5]s", user, pass, address, name, DefaultSSLMode)

	db, err := sql.Open(dbConf.Driver, dsn)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
		return nil
	}
	db.SetMaxOpenConns(MaxOpenConnection)

	dbHandler := &DBHandler{
		Db: db,
	}

	return dbHandler
}
