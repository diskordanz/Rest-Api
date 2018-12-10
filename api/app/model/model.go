package model

import (
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/mock"
)

type Book struct {
	ID   int   	 `gorm:"primary_key" json:"id"`
	Name string  `json:"name"`
	AuthorID int `json:"id_author"`
}

type Author struct {
	ID   int   	 `gorm:"primary_key" json:"id"`
	Name string  `json:"name"`
	Books []Book `json:"books"`
}

type GormDB struct {
	DB *gorm.DB
}

type MockDB struct {
	mock.Mock
}

func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Book{}, &Author{})
	return db
}

type BookService interface {
	GetBooks() ([]Book, error)
	GetBook(book *Book) error
	CreateBook(book *Book) error
	UpdateBook(book *Book) error
	DeleteBook(book *Book) error
	GetFilterBooks(filterString string) error
	GetBooksByAuthor(books *Book) ([]Book, error)
	GetFilterBooksByAuthor(idAuthor int, filterString string) ([]Book, error)

}

type AuthorService interface {
	GetAuthors() ([]Author, error)
	GetAuthor(author *Author) error
	CreateAuthor(author *Author) error
	UpdateAuthor(author *Author) error
	DeleteAuthor(author *Author) error
	GetFilterAuthors(filterString string) error

}

