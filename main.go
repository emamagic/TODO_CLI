package main

import (
	"bufio"
	"flag"
	f "fmt"
	"os"
	"strconv"
	"todo/repository"
	"todo/service"
	"todo/storage"
)

const (
	REGISTER_USER   = "register-user"
	CREATE_TASK     = "create-task"
	CREATE_CATEGORY = "create-category"
	LIST_TASK       = "list-task"
	LOGIN           = "login"
	ME              = "me"
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
	taskRepo := repository.NewTaskRepo()
	categoryRepo := repository.NewCategoryRepo()
	userStore := storage.New(USER_PATH)
	userRepo := repository.NewUserRepo(userStore)

	taskService = service.NewTaksService(taskRepo, categoryRepo, userRepo)
	categoryService = service.NewCategoryService(categoryRepo, userRepo)
	userService = service.NewUserService(userRepo)
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
	case ME:
		AuthUsername()
	case EXIT:
		os.Exit(0)
	}
}

func CreateTask() {
	var title, dueDate string

	f.Println("Please enter a title")
	scanner.Scan()
	title = scanner.Text()

	f.Println("Please enter a dueDate")
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

	_, err := userService.Register(service.RegisterUserRequest{
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

func AuthUsername() {
	authUser, err := userService.UserRepository.AuthUser()
	if err != nil {
		f.Printf("error: %v", err)
	}
	f.Println(authUser.Username, authUser.Password)
}

func getNewCommand(command *string) {
	f.Println("Please enter a command")
	scanner.Scan()
	*command = scanner.Text()
}
