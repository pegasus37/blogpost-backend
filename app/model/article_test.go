package model_test

import (
	"testing"

	"github.com/pegasus37/blogpost-backend/app/model"
	"github.com/stretchr/testify/assert"
)

func TestArticleModel(t *testing.T) {
	article := model.NewArticle()

	t.Run("test_new_article", func(t *testing.T) {
		assert.NotNil(t, article, "article should not be nil")
	})

	t.Run("test_getter", func(t *testing.T) {
		assert.Empty(t, article.Id(), "id should be empty")
		assert.Empty(t, article.Title(), "title should be empty")
		assert.Empty(t, article.Content(), "content should be empty")
		assert.Empty(t, article.Author(), "author should be empty")
		assert.Empty(t, article.CreatedAt(), "createdAt should be empty")
		assert.Empty(t, article.CreatedBy(), "createdBy should be empty")
		assert.Empty(t, article.UpdatedAt(), "updatedAt should be empty")
		assert.Empty(t, article.UpdatedBy(), "updatedBy should be empty")
	})
}
