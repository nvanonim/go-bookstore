package models

import (
	"github.com/nvanonim/go-bookstore/pkg/config"
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
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	result := db.Create(&b)
	if result.Error != nil {
		return nil
	}
	return b
}

func GetAllBooks() []Book {
	var books []Book
	if result := db.Find(&books); result.Error != nil {
		return nil
	}

	return books
}

func GetBookById(id int64) *Book {
	var book Book
	if result := db.First(&book, id); result.Error != nil {
		return nil
	}

	return &book
}

func (b *Book) UpdateBook() *Book {
	result := db.Save(&b)
	if result.Error != nil {
		return nil
	}
	return b
}

func DeleteBookById(id int64) *Book {
	var book Book
	if result := db.First(&book, id); result.Error != nil {
		return nil
	}

	db.Delete(&book)
	return &book
}
