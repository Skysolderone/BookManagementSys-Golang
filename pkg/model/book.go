/*
Create time at 2023年3月9日0009下午 16:22:57
Create User at Administrator
*/

package model

import (
	"github.com/Skysolderone/BookManagementSys-Golang/pkg/config"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

var Db *gorm.DB

func init() {
	config.Connect()
	Db = config.GetDb()
	Db.AutoMigrate(&Book{})
}
func (b *Book) Createbook() *Book {
	Db.Create(&b)
	return b
}
func UpdateBook(book *Book) {
	Db.Save(&book)
}
func GetAllbook() []Book {
	var Books []Book
	Db.Find(&Books)
	return Books
}

func Getbookbyid(id int64) (Book, *gorm.DB) {
	var book Book
	Db.Where("id = ?", id).Find(&book)

	return book, Db
}
func Deletebook(id int64) Book {
	var book Book
	Db.Where("id=?", id).Delete(&book)
	return book
}
