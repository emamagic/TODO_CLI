package repository

import (
	"todo/entity"
)

var categoryStorage []entity.Category

func AddCategory(title, color string) {

	category := entity.Category{
		ID:     len(categoryStorage) + 1,
		Title:  title,
		Color:  color,
		UserID: authenticatedUser.ID,
	}
	categoryStorage = append(categoryStorage, category)
}
