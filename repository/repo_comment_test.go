package repository

import (
	"context"
	"fmt"
	"golang-mysql/database"
	"golang-mysql/entity"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestCommentInsert(t *testing.T) {
	CommentRepository := NewCommentRepository(database.GetConnection())

	ctx := context.Background()

	comment := entity.Comment{
		Email:   "testrepo01@gmail.com",
		Comment: "Test Repo",
	}

	result, err := CommentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestCommentFindById(t *testing.T) {
	CommentRepository := NewCommentRepository(database.GetConnection())

	ctx := context.Background()

	comment, err := CommentRepository.FindById(ctx, 20)
	if err != nil {
		panic(err)
	}

	fmt.Println(comment)

}

func TestCommentFindAll(t *testing.T) {
	CommentRepository := NewCommentRepository(database.GetConnection())

	ctx := context.Background()
	comment, err := CommentRepository.FindAll(ctx)
	if err != nil {
		panic(err)
	}

	for _, value := range comment {
		fmt.Println(value)
	}
}
