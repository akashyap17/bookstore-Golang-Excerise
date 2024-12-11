package routes

import (
	"bookstore/controllers"
	"fmt"
)

func HandleRoute(route string) {
	switch route {
	case "list_books":
		controllers.ListBooks()
	case "add_book":
		err := controllers.CreateBook(1, "Go Programming", "John Doe", 399.99)
		if err != nil {
			fmt.Println("Error", err)
		}
	default:
		fmt.Println("Route not found")
	}
}
