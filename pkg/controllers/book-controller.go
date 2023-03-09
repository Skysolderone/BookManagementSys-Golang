/*
Create time at 2023年3月9日0009下午 16:22:48
Create User at Administrator
*/

package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/Skysolderone/BookManagementSys-Golang/pkg/model"
	"github.com/Skysolderone/BookManagementSys-Golang/pkg/utils"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {
	Createbook := &model.Book{}
	utils.ParseBody(r, Createbook)
	b := Createbook.Createbook()
	res, _ := json.Marshal(b)
	w.Header().Set("content-type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookbyBookId(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookID := params["id"]
	id, err := strconv.ParseInt(bookID, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	book, _ := model.Getbookbyid(id)
	res, _ := json.Marshal(book)
	w.Header().Set("content-type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	if book.Name == "" {
		w.Write([]byte("not found"))
	} else {
		w.Write(res)
	}
}

func UpdateBookByBookId(w http.ResponseWriter, r *http.Request) {
	updatebook := &model.Book{}
	utils.ParseBody(r, updatebook)
	params := mux.Vars(r)
	bookID := params["id"]
	id, err := strconv.ParseInt(bookID, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookdetails, _ := model.Getbookbyid(id)
	if bookdetails.Name != "" {
		bookdetails.Name = updatebook.Name
	}
	if bookdetails.Author != "" {
		bookdetails.Author = updatebook.Author
	}
	if bookdetails.Publication != "" {
		bookdetails.Publication = updatebook.Publication
	}
	model.UpdateBook(&bookdetails)
	res, _ := json.Marshal(bookdetails)
	w.Header().Set("content-type", "pkgliaction/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBookByBookId(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookID := params["id"]
	id, err := strconv.ParseInt(bookID, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	model.Deletebook(id)
	w.Header().Set("content-type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("delete success"))
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	books := model.GetAllbook()
	res, _ := json.Marshal(books)
	w.Header().Set("content-type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
