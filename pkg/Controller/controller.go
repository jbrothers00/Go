package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jbrothers0028/Go/pkg/models"
	"github.com/jbrothers0028/Go/pkg/utils"
)

var NewBooks models.Book

func GetBooks(w http.ResponseWriter, r *http.Request) {
	newBook := models.GetBooks()
	res, _ := json.Marshal(newBook)
	w.Header().Set("Conten-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBooksById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetails, _ = models.GetBooksById(ID)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Conten-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func GetBooks(w http.ResponseWriter, r *http.Request) {
	newBook := models.GetBooks()
	res, _ := json.Marshal(newBook)
	w.Header().Set("Conten-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func GetBooksById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error While Parsing")
	}
	bookDetails, _ := models.GetBooksById(ID)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Conten-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	b := CreateBook.CreateBook()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBooksById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookID := vars["bookID"]
	ID, err := strconv.ParseInt(bookID, 0, 0)
	if err != nil {
		fmt.Println("Error while Parsing")
	}
	book := models.DeleteBook(ID)
	res, _ := json.Marshal(book)
	w.Header().Set("Conten-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func UpdateBookById(w http.ResponseWriter, r *http.Request) {
	var updateBook = &models.Book{}
	utils.ParseBody(r, updateBook)
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error While Parsing")
	}
	bookDetails, db := models.GetBooksById(ID)
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Conten-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
