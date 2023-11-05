package service

import (
	"fmt"
	"todo/entity"
)

type CategoryRepository interface {
	DoesThisUserHaveThisCategoryID(userID, categoryID int) bool
	CreateCategory(cat entity.Category) entity.Category
}

type Category struct {
	categoryRepository CategoryRepository
	userRepository     UserRepository
}

func NewCategoryService(categoryRepo CategoryRepository, userRepo UserRepository) Category {
	return Category{
		categoryRepository: categoryRepo,
		userRepository:     userRepo,
	}
}

type CreateCategoryResponse struct {
	Category entity.Category
	Metadata string
}

type CreateCategoryRequest struct {
	Title string
	Color string
}

func (cat Category) Create(req CreateCategoryRequest) (CreateCategoryResponse, error) {

	authenticatedUser, uErr := cat.userRepository.AuthUser()
	if uErr != nil {
		return CreateCategoryResponse{}, fmt.Errorf("can not create new category: %v", uErr)
	}

	category := cat.categoryRepository.CreateCategory(entity.Category{
		Title:  req.Title,
		Color:  req.Color,
		UserID: authenticatedUser.ID,
	})
	return CreateCategoryResponse{
		Category: category,
	}, nil
}
