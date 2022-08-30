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

func (b *Book) GetBook(id int64) error {
	return db.First(&b, "ID=?", id).Find(&b).Error
}

func (b *Book) DeleteBook(id int64) error {
	return db.First(&b, "ID=?", id).Delete(&b).Error
}

func (b *Book) UpdateBook(id int64, body *Book) error {
	db.First(&b, id)

	err := db.Model(&b).Updates(Book{
		Author:      body.Author,
		Name:        body.Name,
		Publication: body.Publication,
	}).Error
	return err
}
