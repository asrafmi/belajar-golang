package service

import (
	"testing"
	"unit-test/entity"
	"unit-test/repository"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func CleanUpMock() *repository.CategoryRepositoryMock {
	repo := &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
	return repo
}

func TestCategoryService_Get(t *testing.T) {
	t.Run("Category is nil", func(t *testing.T) {
		categoryRepository := CleanUpMock()
		categoryService := CategoryService{Repository: categoryRepository}

		categoryRepository.Mock.On("FindById", "1").Return(nil)

		category, err := categoryService.Get("1")
		assert.Nil(t, category)
		assert.NotNil(t, err)
	})
	t.Run("Category is exists", func(t *testing.T) {
		categoryRepository := CleanUpMock()
		categoryService := CategoryService{Repository: categoryRepository}

		mockCategory := entity.Category{
			Id:   "1",
			Name: "Pencil",
		}
		categoryRepository.Mock.On("FindById", "1").Return(mockCategory)

		category, err := categoryService.Get("1")
		assert.Nil(t, err)
		assert.NotNil(t, category)
		assert.Equal(t, mockCategory, *category, "Category must contains Id = 1 and Name = Pencil")
	})
}
