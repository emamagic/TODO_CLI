package service

import (
	"fmt"
	"todo/entity"
)

type TaskRepository interface {
	CreateNewTask(t entity.Task) (entity.Task, error)
	ListUserTask(userID int) ([]entity.Task, error)
}

type Task struct {
	taskRepository     TaskRepository
	categoryRepository CategoryRepository
	userRepository     UserRepository
}

func NewTaksService(taskRepo TaskRepository, categoryRepo CategoryRepository, userRepo UserRepository) Task {
	return Task{
		taskRepository:     taskRepo,
		categoryRepository: categoryRepo,
		userRepository:     userRepo,
	}
}

type CreateTaskRequest struct {
	Title      string
	DueDate    string
	CategoryID int
}

type CreateTaskResponse struct {
	Task     entity.Task
	Metadata string
}

func (t Task) Create(req CreateTaskRequest) (CreateTaskResponse, error) {

	authenticatedUser, uErr := t.userRepository.GetUser()
	if uErr != nil {
		return CreateTaskResponse{}, fmt.Errorf("can not create new task: %v", uErr)
	}

	if !t.categoryRepository.DoesThisUserHaveThisCategoryID(authenticatedUser.ID, req.CategoryID) {
		return CreateTaskResponse{}, fmt.Errorf("user does not have this category: %d", req.CategoryID)
	}

	createdTask, cErr := t.taskRepository.CreateNewTask(entity.Task{
		Title:      req.Title,
		DueDate:    req.DueDate,
		IsDone:     false,
		UserID:     authenticatedUser.ID,
		CategoryID: req.CategoryID,
	})
	if cErr != nil {
		return CreateTaskResponse{}, fmt.Errorf("can not create new task: %v", cErr)
	}

	return CreateTaskResponse{Task: createdTask}, nil
}

type ListResponse struct {
	Tasks []entity.Task
}

func (t Task) List() (ListResponse, error) {

	authenticatedUser, uErr := t.userRepository.GetUser()
	if uErr != nil {
		return ListResponse{}, fmt.Errorf("can not list tasks: %v", uErr)
	}

	tasks, err := t.taskRepository.ListUserTask(authenticatedUser.ID)
	if err != nil {
		return ListResponse{}, fmt.Errorf("can not list tasks: %v", err)
	}

	return ListResponse{Tasks: tasks}, nil
}
