package model_test

import (
	"testing"

	"github.com/pegasus37/blogpost-backend/app/model"
	"github.com/stretchr/testify/assert"
)

func TestArticleModel(t *testing.T) {
	var article model.Article

	t.Run("test_new_article", func(t *testing.T) {
		article = model.NewArticle()

		assert.NotNil(t, article, "article should not be nil")
	})
}
