package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const API_VERSION = "0.1.1"

/*

 */
type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Isbn   string `json:"isbn"`
	Year   int    `json:"year"`
}

// "Library
var Library = map[string]Book{
	"2343242343241": Book{
		Title:  "First Book",
		Author: "Book Author 1",
		Isbn:   "2343242343241",
		Year:   1979,
	},
	"2343242343242": Book{
		Title:  "Second Book",
		Author: "Book Author 2",
		Isbn:   "2343242343242",
		Year:   1980,
	},
}

//GetBook return a book when given the isbn string
func GetBook(isbn string) (Book, bool) {
	if book, ok := Library[isbn]; ok {
		return book, true

	} else {
		fmt.Println("book not found")
		return Book{}, false
	}
}

//UpdateBook updates book matched by isbn string
func UpdateBook(isbn string, book Book) bool {
	_, exists := Library[isbn]
	if exists {
		Library[isbn] = book
	}
	return exists
}

//DeleteBook delete a record from Lbrary for a book with
//isbn number matching the input
func DeleteBook(isbn string) {
	delete(Library, isbn)
}

func (b *Book) toJSON() []byte {
	result, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}
	return result
}

/*
ListBooks
*/
func ListBooks(res http.ResponseWriter, req *http.Request) {

	if req.Method != http.MethodGet {
		fmt.Println("sorry only GET request allowed")
		res.WriteHeader(http.StatusOK)
		res.Write([]byte("Sorry only GET reuqest allowed"))
		return
	}
	bookList, err := json.Marshal(Library)
	if err != nil {
		panic(err)
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	_, err = res.Write([]byte(bookList))
}

/*

 */
func BookActions(res http.ResponseWriter, req *http.Request) {

	switch req.Method {

	case http.MethodPost:
		book := Book{}
		bBook, err := ioutil.ReadAll(req.Body)
		err = json.Unmarshal(bBook, &book)
		if err != nil {
			panic(err)
		}
		_, exists := GetBook(book.Isbn)
		if exists {
			res.WriteHeader(http.StatusConflict)
			res.Write([]byte("already exist"))
			return
		}
		Library[book.Isbn] = book
		res.WriteHeader(http.StatusCreated)
		res.Header().Add("content-type", "application/json")
		res.Write(bBook)

	case http.MethodGet:
		fmt.Println("GET method called")
		isbn := req.URL.Path[len("/books/"):]
		book, _ := GetBook(isbn)
		fmt.Printf("%v: %v", isbn, isbn)
		fmt.Printf("found book %v \n", string(book.toJSON()))
		res.Header().Set("content-type", "application/json")
		res.WriteHeader(http.StatusOK)
		res.Write(book.toJSON())

	case http.MethodPut:
		jsonBody := Book{}
		isbn := req.URL.Path[len("/books/"):]
		body, err := ioutil.ReadAll(req.Body)

		if err != nil {
			panic(err)
		}
		json.Unmarshal(body, &jsonBody)
		fmt.Printf("req.body : %v", string(jsonBody.toJSON()))
		book, _ := GetBook(isbn)

		book.Title = jsonBody.Title
		fmt.Printf("old/new: %v -> %v \n", book.Title, jsonBody.Title)
		UpdateBook(isbn, jsonBody)
		res.Header().Add("Location", "/books")
		//res.WriteHeader(http.StatusTemporaryRedirect)
		res.WriteHeader(http.StatusCreated)
		//res.Write([]byte("Updated"))
	case http.MethodDelete:
		isbn := req.URL.Path[len("/books/"):]
		DeleteBook(isbn)
		res.WriteHeader(http.StatusCreated)
		res.Write([]byte("deleted book"))

	default:
		res.WriteHeader(http.StatusForbidden)
		res.Write([]byte("Invalid enpdoint"))
	}

}
