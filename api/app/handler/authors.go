package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/diskordanz/Rest-Api/api/app/model"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func GetFilterAuthors(db *gorm.DB, w http.ResponseWriter, r *http.Request, filterString string) {
	name := fmt.Sprintf("%%%s%%", filterString)
	var authors []model.Author
	if err := db.Where("name LIKE ?", name).Find(&authors).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, authors)

}

func GetAuthors(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()
	filterString := v.Get("name")
	var authors []model.Author

	if filterString != "" {
		GetFilterAuthors(db, w, r, filterString)
		return
	} else if err := db.Find(&authors).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, authors)

}

func GetAuthor(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var author model.Author

	if err := db.First(&author, params["id"]).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, author)

}

func CreateAuthor(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var author model.Author
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&author); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&author).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, author)
}

func UpdateAuthor(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var author model.Author
	if err := db.First(&author, params["id"]).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return
	}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&author); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()
	if err := db.Save(&author).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, author)
}

func DeleteAuthor(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var author model.Author
	if err := db.First(&author, params["id"]).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return
	}
	if err := db.Delete(&author, params["id"]).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}

