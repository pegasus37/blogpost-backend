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
