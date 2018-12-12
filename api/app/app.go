package app

import (
	"net/http"

	"github.com/diskordanz/Rest-Api/api/app/handler"
	"github.com/diskordanz/Rest-Api/api/app/model"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type App struct {
	Router *mux.Router
	BookServ model.BookService
	AuthorServ model.AuthorService
}

func (a *App) SetRouters() {
	a.Get("/books", a.GetBooks)
	a.Post("/books", a.CreateBook)
	a.Get("/books/{id}", a.GetBook)
	a.Put("/books/{id}", a.UpdateBook)
	a.Delete("/books/{id}", a.DeleteBook)

	a.Get("/authors", a.GetAuthors)
	a.Post("/authors", a.CreateAuthor)
	a.Get("/authors/{id}", a.GetAuthor)
	a.Put("/authors/{id}", a.UpdateAuthor)
	a.Delete("/authors/{id}", a.DeleteAuthor)

	a.Get("/authors/{id_author}/books", a.GetBooksByAuthor)
	a.Get("/authors/{id_author}/books/{id_book}", a.GetBookByAuthor)
}

func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}

func (a *App) GetBooks(w http.ResponseWriter, r *http.Request) {
	handler.GetBooks(a.BookServ, w, r)
}
func (a *App) CreateBook(w http.ResponseWriter, r *http.Request) {
	handler.CreateBook(a.BookServ, w, r)
}
func (a *App) GetBook(w http.ResponseWriter, r *http.Request) {
	handler.GetBook(a.BookServ, w, r)
}
func (a *App) UpdateBook(w http.ResponseWriter, r *http.Request) {
	handler.UpdateBook(a.BookServ, w, r)
}
func (a *App) DeleteBook(w http.ResponseWriter, r *http.Request) {
	handler.DeleteBook(a.BookServ, w, r)
}

func (a *App) GetAuthors(w http.ResponseWriter, r *http.Request) {
	handler.GetAuthors(a.AuthorServ, w, r)
}
func (a *App) CreateAuthor(w http.ResponseWriter, r *http.Request) {
	handler.CreateAuthor(a.AuthorServ, w, r)
}
func (a *App) GetAuthor(w http.ResponseWriter, r *http.Request) {
	handler.GetAuthor(a.AuthorServ, w, r)
}
func (a *App) UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	handler.UpdateAuthor(a.AuthorServ, w, r)
}
func (a *App) DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	handler.DeleteAuthor(a.AuthorServ, w, r)
}

func (a *App) GetBooksByAuthor(w http.ResponseWriter, r *http.Request) {
	handler.GetBooksByAuthor(a.BookServ, w, r)
}
func (a *App) GetBookByAuthor(w http.ResponseWriter, r *http.Request) {
	handler.GetBookByAuthor(a.BookServ, w, r)
}