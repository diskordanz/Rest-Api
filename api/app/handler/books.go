package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"github.com/diskordanz/Rest-Api/api/app/model"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func GetBooks(bookService BookService, w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()
	filterString := v.Get("name")
	var books []model.Book
	var err error

	if filterString != "" {
		name := fmt.Sprintf("%%%s%%", filterString)
		books, err = bookService.GetFilterBooks(name)
	} else {
		books, err = bookService.GetBooks()
	}
	 if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, books)
}

func GetBook(bookService BookService, w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	book := model.Book{ID:id}

	if err := bookService.GetBook(&book); err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return
	}

	respondJSON(w, http.StatusOK, book)
}

func GetBooksByAuthor(bookService BookService, w http.ResponseWriter, r *http.Request) {
	// params := mux.Vars(r)
	// v := r.URL.Query()
	// filterString := v.Get("name")

	// authorID, err := strconv.Atoi(params["id_author"])
	// if err != nil {
	// 	respondError(w, http.StatusBadRequest, err.Error())
	// 	return
	// }

	// if err := db.First(&author, authorID).Error; err != nil {
	// 	respondError(w, http.StatusNotFound, err.Error())
	// 	return
	// }

	// if filterString != "" {
	// 	GetFilterBooksByAuthor(db, w, r, filterString, authorID)
	// 	return
	// } else if err := db.Model(&author).Related(&books).Error; err != nil {
	// 	respondError(w, http.StatusNotFound, err.Error())
	// 	return
	// }
	// respondJSON(w, http.StatusOK, books)

	v := r.URL.Query()
	filterString := v.Get("name")
	var books []model.Book
	var err error

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	if filterString != "" {
		name := fmt.Sprintf("%%%s%%", filterString)
		books, err = bookService.GetFilterBooksByAuthor(id, name)
	} else {
		book := model.Book{ID:id}
		books, err = bookService.GetBooksByAuthor(book)
	}
	 if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, books)
}

func GetBookByAuthor(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var book model.Book

	authorID, err := strconv.Atoi(params["id_author"])
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	bookID, err := strconv.Atoi(params["id_book"])
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	book := model.Book{ID:id, AuthorID:authorID}

	if err := bookService.GetBook(&book); err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return
	}

	respondJSON(w, http.StatusOK, book)
}


////////////////
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
	if err := db.First(&book, params["id"]).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return
	}
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
