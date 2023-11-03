package main

import (
	"bufio"
	"flag"
	f "fmt"
	"os"
	"todo/handler"
)

const (
	ME              = "me"
	REGISTER_USER   = "register-user"
	CREATE_TASK     = "create-task"
	CREATE_CATEGORY = "create-category"
	LIST_TASK       = "list-task"
	TASK            = "task"
	EXIT            = "exit"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {

	f.Println("Welcom to TODO App")

	command := flag.String("command", "no Command", "command to run")
	flag.Parse()

	for {
		checkForAuthentication(command)

		run(*command)
		
		getNewCommand(command)
	}

}

func run(command string) {

	switch command {
	case ME:
		handler.PrintMe()
	case REGISTER_USER:
		handler.RegisterUser()
	case CREATE_TASK:
		handler.CreateTask()
	case CREATE_CATEGORY:
		handler.CreateCategory()
	case TASK:
		handler.GetTask()
	case LIST_TASK:
		handler.ListTask()
	case EXIT:
		os.Exit(0)
	}
}

func getNewCommand(command *string) {
	f.Println("Please enter a command")
	scanner.Scan()
	*command = scanner.Text()
}

func checkForAuthentication(command *string) {
	if *command != EXIT && *command != REGISTER_USER {
		ok := handler.IsUserAuthorized()
		if !ok {
			handler.LoginProcess()
		}
	}
}
