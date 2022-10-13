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

func (b *Book) GetAllBooks() []Book {
	var books []Book
	if result := db.Find(&books); result.Error != nil {
		return nil
	}

	return books
}

func (b *Book) GetBookById(id int64) *Book {
	if result := db.First(&b, id); result.Error != nil {
		return nil
	}

	return b
}

func (b *Book) UpdateBookById(id int64) *Book {
	if result := db.First(&b, id); result.Error != nil {
		return nil
	}

	db.Save(&b)
	return b
}

func (b *Book) DeleteBookById(id int64) *Book {
	if result := db.First(&b, id); result.Error != nil {
		return nil
	}

	db.Delete(&b)
	return b
}
