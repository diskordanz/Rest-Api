package app

import (
	"fmt"
	"log"
	"net/http"
	"rest-api/api/app/handler"
	"rest-api/api/app/model"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// App has router and db instances
type App struct {
	Router *mux.Router
	DB     *gorm.DB
}
type Config struct {
	DB *DBConfig
}

type DBConfig struct {
	Dialect  string
	Username string
	Password string
	Name     string
	Host     string
	Port     string
}

// Initialize App initialize with predefined configuration
func (a *App) Initialize(config *Config) {

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

// Set all required routers
func (a *App) setRouters() {
	// Routing for handling the projects
	a.Get("/books", a.GetAllBooks)
	a.Post("/books", a.CreateBook)
	a.Get("/books/{id}", a.GetBook)
	a.Put("/books/{id}", a.UpdateBook)
	a.Delete("/books/{id}", a.DeleteBook)
}

// func (a *App) GetFilterBooks(w http.ResponseWriter, r *http.Request) {
// 	handler.GetFilterBooks(a.DB, w, r)
// }

// // Get Wrap the router for GET method
// func (a *App) Find(path string, f func(w http.ResponseWriter, r *http.Request)) {
// 	a.Router.HandleFunc(path, f).Queries("name", "{name}").Methods("GET")
// }

// Get Wrap the router for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

// Post Wrap the router for POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

// Put Wrap the router for PUT method
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

// Delete Wrap the router for DELETE method
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}

// GetAllBooks Handlers to manage Book Data
func (a *App) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	handler.GetAllBooks(a.DB, w, r)
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

// Run the app on it's router
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}
