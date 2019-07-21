package repository

import "github.com/pegasus37/blogpost-backend/app/model"

type ArticelRepository interface {
	Insert(article *model.Article) (*model.Article, error)
}
