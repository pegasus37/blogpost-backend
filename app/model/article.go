package model

// Article contains the blogpost
type Article struct {
	id        int64
	title     string
	content   string
	author    string
	createdAt int64
	createdBy string
	updatedAt int64
	updatedBy string
}

func (a Article) Id() int64 {
	return a.id
}

func (a Article) Title() string {
	return a.title
}

func (a Article) Content() string {
	return a.content
}

func (a Article) Author() string {
	return a.author
}

func (a Article) CreatedAt() int64 {
	return a.createdAt
}

func (a Article) CreatedBy() string {
	return a.createdBy
}

func (a Article) UpdatedAt() int64 {
	return a.updatedAt
}

func (a Article) UpdatedBy() string {
	return a.updatedBy
}

func NewArticle() *Article {
	return &Article{}
}
