package model

import "github.com/jinzhu/gorm"

// Book struct
type Book struct {
	ID   uint   `gorm:"primary_key" json:"id"`
	Name string `json:"name"`
}

//
// DBMigrate will create and migrate the tables, and then make the some relationships if necessary
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Book{})
	return db
}
