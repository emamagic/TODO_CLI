package repository

import "todo/entity"

type Task struct {
	tasks []entity.Task
}

func NewTaskRepo() *Task {
	return &Task {
		tasks: make([]entity.Task, 0),
	}
}

func (t *Task) CreateTask(task entity.Task) (entity.Task, error) {
	task.ID = len(t.tasks) + 1
	t.tasks = append(t.tasks, task)

	return task, nil
}

func (task *Task) ListUserTask(userID int) ([]entity.Task, error) {
	var tasks []entity.Task
	for _, t := range task.tasks {
		if t.UserID == userID {
			tasks = append(tasks, t)
		}
	}

	return tasks, nil
}