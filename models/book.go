package models

type Book struct {
	ID     int
	Title  string
	Author string
	Price  float64
}

var books []Book

// Addbook adds a book to the list.
func AddBook(b Book) {
	books = append(books, b)
}

// GEtBooks retrieves all books.
func GetBooks() []Book {
	return books
}

// FindBookByID searches for a book by its ID.

func FindBookByID(id int) *Book {
	for _, b := range books {
		if b.ID == id {
			return &b
		}
	}
	return nil
}
