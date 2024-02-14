// Main Process of the data

package controllers

import (
	"fmt"
	"net/http"

	"github.com/Kaushal_26/go-bookstore/pkg/models"
	"github.com/Kaushal_26/go-bookstore/pkg/utils"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET : Request for All Books\n")

	newBooks := models.GetAllBooks()
	utils.PrintBook(w, newBooks)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	fmt.Print("GET : Request for Book By ID = ")

	ID := utils.BookIdFind(r)
	fmt.Println(ID, "\n")

	bookDetails, _ := models.GetBookById(ID)
	utils.PrintBook(w, bookDetails)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("POST : Request for Book Creation\n")

	createBook := &models.Book{}
	utils.ParseBody(r, createBook)
	b := createBook.CreateBook()
	utils.PrintBook(w, b)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DELETE : Request for Book deletion\n")

	ID := utils.BookIdFind(r)
	book := models.DeleteBook(ID)
	utils.PrintBook(w, book)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("PUT: request for Book Update\n")

	updateBook := &models.Book{}
	utils.ParseBody(r, updateBook)

	ID := utils.BookIdFind(r)
	bookDetails, db := models.GetBookById(ID)
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publications != "" {
		bookDetails.Publications = updateBook.Publications
	}
	db.Save(&bookDetails)

	utils.PrintBook(w, bookDetails)
}
