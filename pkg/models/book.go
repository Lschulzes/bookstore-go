package models

import (
	"github.com/lschulzes/bookstore-go/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	err := db.AutoMigrate(&Book{})
	if err != nil {
		return
	}
}

func (b *Book) CreateBook() *Book {
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var books []Book
	db.Find(&books)
	return books
}

func (b *Book) GetBook(id int64) (*Book, *gorm.DB) {
	db := db.Where("ID=?", id).Find(&b)

	return b, db
}

func (b *Book) DeleteBook(id int64) error {
	return db.Where("ID=?", id).Delete(&b).Error
}

func (b *Book) UpdateBook(id int64) error {
	err := db.Where("ID=?", id).Updates(&b).Error
	b.ID = uint(id)
	return err
}
