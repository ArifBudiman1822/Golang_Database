package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang-mysql/entity"
	"strconv"
)

type CommentRepositoryImpl struct {
	DB *sql.DB
}

func NewCommentRepository(db *sql.DB) CommentRepository {
	return &CommentRepositoryImpl{DB: db}
}

func (repo *CommentRepositoryImpl) Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	script := "insert into comment(email, comment) values(?, ?)"
	result, err := repo.DB.ExecContext(ctx, script, comment.Email, comment.Comment)
	if err != nil {
		return comment, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return comment, err
	}
	comment.Id = int32(id)
	return comment, nil

}

func (repo *CommentRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Comment, error) {
	script := "select email, comment from comment where id = ? limit 1"
	rows, err := repo.DB.QueryContext(ctx, script, id)
	comment := entity.Comment{}
	if err != nil {
		return comment, err
	}
	defer rows.Close()

	if rows.Next() {
		rows.Scan(&comment.Email, &comment.Comment, &comment.Id)
		return comment, nil
	} else {
		return comment, errors.New("ID" + strconv.Itoa(int(id)) + "Not Found")
	}

}

func (repo *CommentRepositoryImpl) FindAll(ctx context.Context) ([]entity.Comment, error) {
	script := "select email, comment from comment"
	rows, err := repo.DB.QueryContext(ctx, script)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var comments []entity.Comment
	for rows.Next() {
		comment := entity.Comment{}
		rows.Scan(&comment.Email, &comment.Comment, &comment.Id)
		comments = append(comments, comment)
	}

	return comments, nil

}
