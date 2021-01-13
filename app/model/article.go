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

// NewArticle initiate new article model
func NewArticle() *Article {
	return &Article{}
}

// Id get id of article
func (a Article) Id() int64 {
	return a.id
}

// Title get title of article
func (a Article) Title() string {
	return a.title
}

// SetTitle set title of article
func (a *Article) SetTitle(title string) {
	a.title = title
}

// Content get content of article
func (a Article) Content() string {
	return a.content
}

// SetContent set content of article
func (a *Article) SetContent(content string) {
	a.content = content
}

// Author get author of article
func (a Article) Author() string {
	return a.author
}

// SetAuthor set author of article
func (a *Article) SetAuthor(author string) {
	a.author = author
}

// CreatedAt get created time of article
func (a Article) CreatedAt() int64 {
	return a.createdAt
}

// CreatedBy get creator of article
func (a Article) CreatedBy() string {
	return a.createdBy
}

// SetCreatedBy set creator of article
func (a *Article) SetCreatedBy(createdBy string) {
	a.createdBy = createdBy
}

// UpdatedAt get update time of article
func (a Article) UpdatedAt() int64 {
	return a.updatedAt
}

// UpdatedBy get updater of article
func (a Article) UpdatedBy() string {
	return a.updatedBy
}

// SetUpdatedBy set updater of article
func (a *Article) SetUpdatedBy(updatedBy string) {
	a.updatedBy = updatedBy
}
