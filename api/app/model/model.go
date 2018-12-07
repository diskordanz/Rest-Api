package model

import "github.com/jinzhu/gorm"

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

func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Book{}, &Author{})
	return db
}

