// "Library
package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBook(t *testing.T) {
	isbn := "2343242343241"
	book, exists := GetBook(isbn)
	if exists {
		assert.Equal(t, book, Library[isbn])
	}
}

func TestUpdateBook(t *testing.T) {
	isbn := "2343242343241"
	book := Book{
		Title:  "Test Title Update",
		Author: "Book Author 1",
		Isbn:   "2343242343241",
		Year:   1979,
	}
	exists := UpdateBook(isbn, book)
	if exists {
		assert.Equal(t, book.Title, Library[isbn].Title)
	}
}

func TestDeleteBook(t *testing.T) {
	DeleteBook("2343242343241")
	_, exists := GetBook("2343242343241")
	assert.Equal(t, len(Library), 1)
	assert.Equal(t, exists, false)
}
