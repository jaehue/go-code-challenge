package models

import (
	"context"
	"errors"
)

func GetProductsById(ctx context.Context, productIds []int64) ([]Product, error) {
	if len(productIds) == 0 {
		return nil, nil
	}

	var products []Product
	if err := DB(ctx).In("id", productIds).Desc("id").Find(&products); err != nil {
		return nil, err
	}

	return products, nil
}

func (b *Brand) Create(ctx context.Context) error {
	if exist, err := DB(ctx).Where("name = ?", b.Name).Exist(&Brand{}); err != nil {
		return err
	} else if exist {
		return errors.New("AlreadyExist")
	}

	if _, err := DB(ctx).Insert(b); err != nil {
		return err
	}

	return nil
}

func (p *Product) Create(ctx context.Context) error {
	if _, err := DB(ctx).Insert(p); err != nil {
		return err
	}

	return nil
}

type Brand struct {
	Id   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type Product struct {
	Id      int64  `json:"id"`
	BrandId int64  `json:"brandId" xorm:"index"`
	Name    string `json:"name" xorm:"index"`
}
