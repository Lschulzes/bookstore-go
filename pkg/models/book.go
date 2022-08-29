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
