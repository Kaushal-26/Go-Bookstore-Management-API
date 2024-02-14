// Create the structure of the database

package models

import (
	"github.com/Kaushal_26/go-bookstore/pkg/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Database
var db *gorm.DB

// What to store
type Book struct {
	gorm.Model
	Name         string `gorm:""json":"name"`
	Author       string `json:"author"`
	Publications string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

// Working with GORM

// Method to create new book
func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

// Get All Books
func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

// Get Book by its id
func GetBookById(ID int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", ID).Find(&getBook)
	return &getBook, db
}

// Delete Book by Id
func DeleteBook(ID int64) Book {
	var book Book
	db.Where("ID=?", ID).Delete(book)
	return book
}
