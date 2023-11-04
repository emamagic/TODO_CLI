package main

import (
	"bufio"
	"flag"
	f "fmt"
	"os"
	"strconv"
	"todo/repository/filestore"
	"todo/repository/memorystore"
	"todo/service"
)

const (
	REGISTER_USER   = "register-user"
	CREATE_TASK     = "create-task"
	CREATE_CATEGORY = "create-category"
	LIST_TASK       = "list-task"
	LOGIN           = "login"
	EXIT            = "exit"

	USER_PATH = "user.txt"
)

var (
	scanner         = bufio.NewScanner(os.Stdin)
	taskService     service.Task
	categoryService service.Category
	userService     service.User
)

func init() {
	taskMemoryRepo := memorystore.NewTaskStore()
	categoryMemoryRepo := memorystore.NewCategoryStore()
	userFileRepo := filestore.NewUserStore(USER_PATH)

	taskService = service.NewTaksService(taskMemoryRepo, categoryMemoryRepo, userFileRepo)
	categoryService = service.NewCategoryService(categoryMemoryRepo, userFileRepo)
	userService = service.NewUserService(userFileRepo)

}

func main() {

	f.Println("Welcom to TODO App")

	command := flag.String("command", "no Command", "command to run")
	flag.Parse()

	for {
		run(*command)

		getNewCommand(command)
	}

}

func run(command string) {

	switch command {
	case REGISTER_USER:
		RegisterUser()
	case CREATE_TASK:
		CreateTask()
	case CREATE_CATEGORY:
		CreateCategory()
	case LIST_TASK:
		ListTask()
	case LOGIN:
		Login()
	case EXIT:
		os.Exit(0)
	}
}

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

	_, cErr := taskService.Create(service.CreateTaskRequest{
		Title:      title,
		DueDate:    dueDate,
		CategoryID: categoryID,
	})
	if cErr != nil {
		f.Printf("error: %v\n", cErr)
	}
}

func ListTask() {
	tasks, err := taskService.List()
	if err != nil {
		f.Printf("error: %v\n", err)
	}
	f.Printf("tasks: %+v\n", tasks)
}

func CreateCategory() {
	var title, color string

	f.Println("Please enter a title")
	scanner.Scan()
	title = scanner.Text()

	f.Println("Please enter a color")
	scanner.Scan()
	color = scanner.Text()

	_, err := categoryService.Create(service.CreateCategoryRequest{
		Title: title,
		Color: color,
	})

	if err != nil {
		f.Printf("error: %v\n", err)
	}
}

func RegisterUser() {
	var username, password string

	f.Println("Please enter a username")
	scanner.Scan()
	username = scanner.Text()

	f.Println("Please enter a password")
	scanner.Scan()
	password = scanner.Text()

	_, err := userService.Create(service.CreateUserRequest{
		Username: username,
		Password: password,
	})
	if err != nil {
		f.Printf("error: %v\n", err)
	}
}

func Login() {
	var username, password string

	f.Println("Login process...")

	f.Println("Please enter a username")
	scanner.Scan()
	username = scanner.Text()

	f.Println("Please enter a password")
	scanner.Scan()
	password = scanner.Text()

	err := userService.Login(service.LoginUserRequest{
		Username: username,
		Password: password,
	})
	if err != nil {
		f.Printf("error: %v\n", err)
	}
}

func getNewCommand(command *string) {
	f.Println("Please enter a command")
	scanner.Scan()
	*command = scanner.Text()
}
