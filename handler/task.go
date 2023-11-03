package handler

import (
	f "fmt"
	"strconv"
	rep "todo/repository"
)

func CreateTask() {
	var title, dueDate string

	f.Println("Please enter a title")
	scanner.Scan()
	title = scanner.Text()

	f.Println("Please enter a duDate")
	scanner.Scan()
	dueDate = scanner.Text()

	f.Println("Please enter the categoryID")
	scanner.Scan()

	categoryID, err := strconv.Atoi(scanner.Text())
	if err != nil {
		f.Printf("category id is not valid integer, %v\n", err)

		return
	}

	rep.AddTask(title, dueDate, categoryID)
}

func ListTask() {
	tasks := rep.LoadTasks()
	for _, t := range tasks {
		f.Printf("task: %+v\n", t)
	}
}

func GetTask() {

	f.Println("Please enter the taskID")
	scanner.Scan()
	taskID, tErr := strconv.Atoi(scanner.Text())
	if tErr != nil {
		f.Printf("task id is not valid integer, %v\n", tErr)

		return
	}

	f.Println("Please enter the categoryID")
	scanner.Scan()
	categoryID, cErr := strconv.Atoi(scanner.Text())
	if cErr != nil {
		f.Printf("category id is not valid integer, %v\n", cErr)

		return
	}

	task := rep.GetTask(taskID, categoryID)
	if task.ID == 0 {
		f.Println("TaskID or CategoryID is not valid")

		return
	}
	f.Printf("task: %+v\n", task)
}
