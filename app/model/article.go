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

func NewArticle() *Article {
	return &Article{}
}

func (a Article) Id() int64 {
	return a.id
}

func (a Article) Title() string {
	return a.title
}

func (a *Article) SetTitle(title string) {
	a.title = title
}

func (a Article) Content() string {
	return a.content
}

func (a *Article) SetContent(content string) {
	a.content = content
}

func (a Article) Author() string {
	return a.author
}

func (a *Article) SetAuthor(author string) {
	a.author = author
}

func (a Article) CreatedAt() int64 {
	return a.createdAt
}

func (a Article) CreatedBy() string {
	return a.createdBy
}

func (a *Article) SetCreatedBy(createdBy string) {
	a.createdBy = createdBy
}

func (a Article) UpdatedAt() int64 {
	return a.updatedAt
}

func (a Article) UpdatedBy() string {
	return a.updatedBy
}

func (a *Article) SetUpdatedBy(updatedBy string) {
	a.updatedBy = updatedBy
}
