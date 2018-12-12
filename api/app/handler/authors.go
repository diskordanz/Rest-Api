package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"github.com/diskordanz/Rest-Api/api/app/model"

	"github.com/gorilla/mux"
)

func GetAuthors(authorService model.AuthorService, w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()
	filterString := v.Get("name")
	var authors []model.Author
	var err error

	if filterString != "" {
		name := fmt.Sprintf("%%%s%%", filterString)
		authors, err = authorService.GetFilterAuthors(name)
	} else {
		authors, err = authorService.GetAuthors()
	}
	 if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, authors)

}

func GetAuthor(authorService model.AuthorService, w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	author := model.Author{ID:id}

	if err := authorService.GetAuthor(&author); err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return
	}

	respondJSON(w, http.StatusOK, author)

}

func CreateAuthor(authorService model.AuthorService, w http.ResponseWriter, r *http.Request) {
	var author model.Author
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&author); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := authorService.CreateAuthor(&author); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, author)
}

func UpdateAuthor(authorService model.AuthorService, w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	author := model.Author{ID:id}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&author); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()
	if err := authorService.UpdateAuthor(&author); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, author)
}

func DeleteAuthor(authorService model.AuthorService, w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	author := model.Author{ID:id}

	if err := authorService.DeleteAuthor(&author); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}

