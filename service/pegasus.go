package service

import (
	"fmt"

	"github.com/pegasus37/blogpost-backend/config"
	"github.com/pegasus37/blogpost-backend/entity"
	"github.com/pegasus37/blogpost-backend/repository"
)

func GetPegasus(config *config.Configuration) *entity.Pegasus {
	fmt.Println("[LOG: Service: GetPegasus, Message: Get Pegasus]")

	postgreRepo := repository.NewPostgresRepo(config.DB)
	// mySqlRepo := repository.NewMySqlRepo(config.DB)

	repo := repository.NewRepository(postgreRepo)

	pegasus, err := repo.Repo.Get()

	if err != nil {
		fmt.Println("[LOG - ERROR: Service: GetPegasus, Error: ", err)
	}

	return &pegasus
}

func InsertPegasus(config *config.Configuration, pegasus []entity.Pegasus) error {
	fmt.Println("[LOG: Service: InsertPegasus, Message: Insert Pegasus]")

	postgreRepo := repository.NewPostgresRepo(config.DB)

	repo := repository.NewRepository(postgreRepo)

	err := repo.Repo.Insert(pegasus)

	if err != nil {
		fmt.Println("[LOG - ERROR: Service: GetPegasus, Error: ", err)
	}

	return nil
}
