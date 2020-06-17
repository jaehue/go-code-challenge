package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPost(t *testing.T) {
	testData := struct {
		brand    string
		products []string
		posts    []string
	}{
		brand:    "brand#1",
		products: []string{"product#1", "product#2", "product#3", "product#4", "product#5"},
		posts:    []string{"post#1", "post#2", "post#3", "post#4"},
	}

	err := (&Brand{Name: testData.brand}).Create(ctx)
	assert.Nil(t, err)

	for _, productName := range testData.products {
		err := (&Product{BrandId: 1, Name: productName}).Create(ctx)
		assert.Nil(t, err)
	}

	for _, postTitle := range testData.posts {
		err := (&Post{Title: postTitle}).Create(ctx)
		assert.Nil(t, err)
	}

	t.Run("AddProductToPost", func(t *testing.T) {
		products, err := AddProductToPost(ctx, 1, []int64{1, 3, 3, 5})
		assert.Nil(t, err)
		assert.Equal(t, len(products), 3)
		assert.EqualValues(t, products[0].Id, 5)
		assert.EqualValues(t, products[1].Id, 3)
		assert.EqualValues(t, products[2].Id, 1)
	})
	t.Run("RemoveProductFromPost", func(t *testing.T) {
		err := RemoveProductFromPost(ctx, 1, []int64{1, 1, 5})
		assert.Nil(t, err)
		post, err := GetPost(ctx, 1, 10)
		assert.Nil(t, err)
		assert.Equal(t, len(post.Products), 1)
		assert.EqualValues(t, post.Products[0].Id, 3)
	})

}
