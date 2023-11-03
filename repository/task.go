package repository

import (
	f "fmt"
	"todo/entity"
)

var taskStorage []entity.Task

func LoadTasks() []entity.Task {
	var tasks []entity.Task

	for _, t := range taskStorage {
		if authenticatedUser.ID == t.UserID {
			tasks = append(tasks, t)
		}
	}

	return tasks
}

func AddTask(title, dueData string, categoryID int) {

	isFound := false
	for _, c := range categoryStorage {
		if categoryID == c.ID {
			isFound = true
		}
	}
	if !isFound {
		f.Printf("CategoryID is not valid %d\n", categoryID)

		return
	}

	task := entity.Task{
		ID:         len(taskStorage) + 1,
		Title:      title,
		DueDate:    dueData,
		IsDone:     false,
		UserID:     authenticatedUser.ID,
		CategoryID: categoryID,
	}

	taskStorage = append(taskStorage, task)
}

func GetTask(taskID, categoryID int) entity.Task {

	for _, t := range taskStorage {
		if authenticatedUser.ID == t.UserID && taskID == t.ID && categoryID == t.CategoryID {
			return t
		}
	}

	return entity.Task{}
}
