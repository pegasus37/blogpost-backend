package repository

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	_ "github.com/lib/pq"
	"gitlab.mapan.io/scn/sc_web/module/sc_schema/exception"

	"github.com/pegasus37/blogpost-backend/config"
	"github.com/pegasus37/blogpost-backend/entity"
)

const (
	DefaultSSLMode    = "disable"
	MaxOpenConnection = 40
)

type PostgresRepo struct {
	Db *sql.DB
}

func NewPostgresRepo(config *config.DBConfig) *PostgresRepo {
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

	postgresRepo := &PostgresRepo{
		Db: db,
	}

	return postgresRepo
}

func (p PostgresRepo) Get() (entity.Pegasus, error) {

	item := entity.Pegasus{}
	query := "SELECT name, number, type, year FROM pegasus"

	err := p.Db.QueryRow(fmt.Sprintf(query)).Scan(&item.Name, &item.Number, &item.Type, &item.Year)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("[LOG: Repo: Get, Message: No Rows]")
			return item, exception.NewNotFoundException("Not Found Exception")
		} else {
			return item, err
		}
	}

	return item, nil
}

func (p PostgresRepo) Insert(items []entity.Pegasus) error {
	tx, err := p.Db.Begin()

	if err != nil {
		return err
	}

	for _, item := range items {
		query := "INSERT INTO pegasus(name, number, type, year) VALUES($1, $2, $3, $4)"

		params := []interface{}{
			item.Name,
			item.Number,
			item.Type,
			item.Year,
		}

		stmt, err := p.Db.Prepare(query)

		if err != nil {
			tx.Rollback()
			return err
		}

		defer stmt.Close()

		if _, err := stmt.Exec(params...); err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()

	return nil
}
