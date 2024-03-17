package models

import (
	"github.com/jbrothers0028/Go/pkg/Config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:""json:"name"`
	Author      string `gorm:""json:"author"`
	Publication string `gorm:""json:"publication"`
}

func init() {
	Config.Connect()
	db = Config.getDB()
	db.AutoMigrate(&Book{})
}
func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}
func GetBooksById(Id int64) (*Book, *gorm.DB) {
	var getbook Book
	db.Where("ID=?", Id).Find(&getbook)
	return &getbook, db
}
func DeleteBooksById(Id int64) Book {
	var book Book
	db.Where("ID=?", Id).Delete(&book)
	return book
}
