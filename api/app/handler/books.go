package handler

import (
	"encoding/json"
	"net/http"
	"rest-api/api/app/model"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func GetAllBooks(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var books []model.Book
	db.Select("name").Find(&books)
	respondJSON(w, http.StatusOK, &books)
}

func GetBook(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var book model.Book

	if err := db.First(&book, params["id"]).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, book)

}

func CreateBook(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var book model.Book
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&book); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&book).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, book)
}

func UpdateBook(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var book model.Book
	if err := db.First(&book, params["id"]).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return
	}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&book); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()
	if err := db.Save(&book).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, book)
}

func DeleteBook(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var book model.Book

	if err := db.Delete(&book, params["id"]).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}

// getBookOr404 gets a book instance if exists, or respond the 404 error otherwise
// func getBookOr404(db *gorm.DB, name string, w http.ResponseWriter, r *http.Request) *model.Book {
// 	book := model.Book{}
// 	if err := db.First(&book, model.Book{Name: name}).Error; err != nil {
// 		respondError(w, http.StatusNotFound, err.Error())
// 		return nil
// 	}
// 	return &book
// }