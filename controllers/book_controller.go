package controllers

import (
	"bookstore/models"
	"errors"
	"fmt"
)

func CreateBook(id int, title, author string, price float64) error {
	if models.FindBookByID(id) != nil {
		return errors.New("Book with the same ID already exists")
	}
	models.AddBook(models.Book{ID: id, Title: title, Author: author, Price: price})
	return nil
}

func ListBooks() {
	for _, book := range models.GetBooks() {
		fmt.Printf("ID: %d, Title: %s, Author: %s, Price: %.2f\n", book.ID, book.Title, book.Author, book.Price)
	}
}
