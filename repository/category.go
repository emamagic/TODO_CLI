package repository

import "todo/entity"

type Category struct {
	categories []entity.Category
}

func NewCategoryRepo() *Category {
	return &Category{
		categories: make([]entity.Category, 0),
	}
}

func (cat Category) DoesThisUserHaveThisCategoryID(userID, categoryID int) bool {
	isFound := false
	for _, c := range cat.categories {
		if categoryID == c.ID && userID == c.UserID {
			isFound = true

			break
		}
	}

	return isFound
}

func (c *Category) CreateCategory(cat entity.Category) entity.Category {
	cat.ID = len(c.categories) + 1
	c.categories = append(c.categories, cat)
	return cat
}
