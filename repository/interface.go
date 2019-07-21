package repository

import "github.com/pegasus37/blogpost-backend/entity"

type RepositoryInterface interface {
	Get() (entity.Pegasus, error)
	Insert(items []entity.Pegasus) error
}
