package models

import (
	"context"
	"errors"
	"time"
)

func GetPost(ctx context.Context, postId int64, subItemCount int) (*Post, error) {
	var p Post
	if exist, err := DB(ctx).ID(postId).Get(&p); err != nil {
		return nil, err
	} else if !exist {
		return nil, errors.New("ErrorNotFound")
	}

	if err := DB(ctx).
		Join("INNER", "post_product", "product.id = post_product.product_id").
		Where("post_product.post_id = ?", p.Id).
		Limit(subItemCount).
		Find(&p.Products); err != nil {
		return nil, err
	}

	return &p, nil
}
func (p *Post) Create(ctx context.Context) error {
	if _, err := DB(ctx).Insert(p); err != nil {
		return err
	}

	return nil
}
func RemoveProductFromPost(ctx context.Context, postId int64, productIds []int64) error {
	return nil
}

func AddProductToPost(ctx context.Context, postId int64, productIds []int64) ([]Product, error) {
	return nil, nil
}

type Post struct {
	Id    int64  `json:"id"`
	Title string `json:"title" xorm:"index"`

	Products []Product `json:"products" xorm:"-"`
}

type PostProduct struct {
	Id        int64     `json:"id"`
	PostId    int64     `json:"postId" xorm:"index"`
	ProductId int64     `json:"productId" xorm:"index"`
	CreatedAt time.Time `json:"createdAt" xorm:"created"`
}
