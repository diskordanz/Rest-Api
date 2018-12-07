package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/diskordanz/Rest-Api/api/app/handler"
	"github.com/diskordanz/Rest-Api/api/app/model"
	"github.com/diskordanz/Rest-Api/api/config"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

func (a *App) Initialize(config *config.Config) {

	dbURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		config.DB.Host,
		config.DB.Port,
		config.DB.Username,
		config.DB.Name,
		config.DB.Password)

	db, err := gorm.Open(config.DB.Dialect, dbURI)
	if err != nil {
		log.Fatal(err)
	}

	a.DB = model.DBMigrate(db)
	a.Router = mux.NewRouter()
	a.setRouters()
}

func (a *App) setRouters() {
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

func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
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
	handler.GetBooks(a.DB, w, r)
}
func (a *App) CreateBook(w http.ResponseWriter, r *http.Request) {
	handler.CreateBook(a.DB, w, r)
}
func (a *App) GetBook(w http.ResponseWriter, r *http.Request) {
	handler.GetBook(a.DB, w, r)
}
func (a *App) UpdateBook(w http.ResponseWriter, r *http.Request) {
	handler.UpdateBook(a.DB, w, r)
}
func (a *App) DeleteBook(w http.ResponseWriter, r *http.Request) {
	handler.DeleteBook(a.DB, w, r)
}


func (a *App) GetAuthors(w http.ResponseWriter, r *http.Request) {
	handler.GetAuthors(a.DB, w, r)
}
func (a *App) CreateAuthor(w http.ResponseWriter, r *http.Request) {
	handler.CreateAuthor(a.DB, w, r)
}
func (a *App) GetAuthor(w http.ResponseWriter, r *http.Request) {
	handler.GetAuthor(a.DB, w, r)
}
func (a *App) UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	handler.UpdateAuthor(a.DB, w, r)
}
func (a *App) DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	handler.DeleteAuthor(a.DB, w, r)
}

func (a *App) GetBooksByAuthor(w http.ResponseWriter, r *http.Request) {
	handler.GetBooksByAuthor(a.DB, w, r)
}
func (a *App) GetBookByAuthor(w http.ResponseWriter, r *http.Request) {
	handler.GetBookByAuthor(a.DB, w, r)
}