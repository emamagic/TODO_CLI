package handler

import (
	"bufio"
	f "fmt"
	"os"
	rep "todo/repository"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
)

func RegisterUser() {
	var username, pass string

	f.Println("Please set your username")
	scanner.Scan()
	username = scanner.Text()

	f.Println("Please set your pass")
	scanner.Scan()
	pass = scanner.Text()

	rep.AddUser(username, pass)
}

func LoginProcess() {
	var username, pass string

	f.Println("Login Process...")

	f.Println("Please enter your username")
	scanner.Scan()
	username = scanner.Text()

	f.Println("Please enter your pass")
	scanner.Scan()
	pass = scanner.Text()

	ok := rep.LoadUser(username, pass)
	if !ok {
		f.Println("You have to register fist")
		RegisterUser()
	}
}

func PrintMe() {
	f.Printf("you: %s\n", rep.CurrentUser())
}

func IsUserAuthorized() bool {
	return rep.CurrentUser() != "" 
}

