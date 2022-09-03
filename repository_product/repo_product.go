package repository_product

import (
	"context"
	"golang-mysql/entity"
)

type ProductRepository interface {
	Insert(ctx context.Context, product entity.Product) (entity.Product, error)
	FindById(ctx context.Context, id int) (entity.Product, error)
	FindAll(ctx context.Context) (entity.Product, error)
}
