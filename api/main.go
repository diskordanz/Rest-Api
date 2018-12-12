package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/diskordanz/Rest-Api/api/app"
	"github.com/diskordanz/Rest-Api/api/config"
	"github.com/diskordanz/Rest-Api/api/app/model"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

var applic *app.App

func main() {
	log.Fatal(http.ListenAndServe((":8080"), applic.Router))
}

func init() {
	c := config.GetConfig()
	dbURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		c.DB.Host,
		c.DB.Port,
		c.DB.Username,
		c.DB.Name,
		c.DB.Password)

	db, err := gorm.Open(c.DB.Dialect, dbURI)
	if err != nil {
		log.Fatal(err)
	}

	applic = &app.App{Router: mux.NewRouter(), BookServ: &model.GormDB{DB: db}, AuthorServ: &model.GormDB{DB: db}}	
	applic.SetRouters()
}
