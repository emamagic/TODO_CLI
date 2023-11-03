package handler

import (
	f "fmt"
	rep "todo/repository"
)

func CreateCategory() {
	var title, color string

	f.Println("Please enter a title")
	scanner.Scan()
	title = scanner.Text()

	f.Println("Please enter a color")
	scanner.Scan()
	color = scanner.Text()

	rep.AddCategory(title, color)
}
