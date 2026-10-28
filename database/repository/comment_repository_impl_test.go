package repository

import (
	"context"
	"database"
	"database/entity"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestCommentInsert(t *testing.T) {
	commentRepo := NewCommentRepository(database.GetConnection())

	ctx := context.Background()
	comment := entity.Comment{
		Email:   "test_repo@test.com",
		Comment: "testing repo",
	}

	result, err := commentRepo.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestFindById(t *testing.T) {
	commentRepo := NewCommentRepository(database.GetConnection())

	ctx := context.Background()

	result, err := commentRepo.FindById(ctx, 32)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestFindAll(t *testing.T) {
	commentRepo := NewCommentRepository(database.GetConnection())

	ctx := context.Background()

	result, err := commentRepo.FindAll(ctx)
	if err != nil {
		panic(err)
	}

	for _, comment := range result {
		fmt.Println(comment)
	}
}
