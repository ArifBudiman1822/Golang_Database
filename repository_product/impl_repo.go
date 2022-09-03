package repository_product

import (
	"context"
	"database/sql"
	"errors"
	"golang-mysql/entity"
	"strconv"
)

type ProductRepositoryImpl struct {
	DB *sql.DB
}

func (repo *ProductRepositoryImpl) Insert(ctx context.Context, product entity.Product) (entity.Product, error) {
	script := "INSERT INTO products(products,quantity,price)VALUES(?, ?, ?)"
	result, err := repo.DB.ExecContext(ctx, script)
	if err != nil {
		return product, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return product, err
	}
	product.Id = (int(id))
	return product, nil
}

func (repo *ProductRepositoryImpl) FindById(ctx context.Context, id int) (entity.Product, error) {
	script := "select product, quantity, price from product where id = ? limit 1"
	product := entity.Product{}
	rows, err := repo.DB.QueryContext(ctx, script, id)
	if err != nil {
		return product, err
	}
	defer rows.Close()

	if rows.Next() {
		rows.Scan(&product.Id, &product.Product, &product.Quantity, &product.Price)
		return product, nil
	} else {
		return product, errors.New("ID" + strconv.Itoa(id) + "Not Found")
	}
}

func (repo *ProductRepositoryImpl) FindAll(ctx context.Context) ([]entity.Product, error) {
	script := "select id, product,quantity,price from products"
	var products []entity.Product
	rows, err := repo.DB.QueryContext(ctx, script)
	if err != nil {
		return products, err
	}
	defer rows.Close()

	for rows.Next() {
		product := entity.Product{}
		rows.Scan(&product.Id, &product.Product, &product.Quantity, &product.Price)
		products = append(products, product)
	}
	return products, nil
}
