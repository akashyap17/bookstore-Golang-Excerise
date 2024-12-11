//package main
//
//import (
//	"bookstore/routes"
//	"bookstore/services"
//)
//
//func main() {
//	// Simulate adding and listing books
//	routes.HandleRoute("add_book")
//	routes.HandleRoute("list_books")
//
//	// Simulate sending notifications
//	services.NotifyAllCustomers()
//
//	// Save logs to a file
//	services.SaveToFile("logs.txt", "Bookstore initialized successfully")
//}

package main

import (
	"bookstore/routes"
	"bookstore/services"
)

func main() {
	// Simulate adding and listing books
	routes.HandleRoute("add_book")
	routes.HandleRoute("list_books")

	// Simulate sending notifications
	services.NotifyAllCustomers()

	// Save logs to a file
	services.SaveToFile("logs.txt", "Bookstore initialized successfully.")
}
