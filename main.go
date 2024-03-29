package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Kei-K23/go-url-shortened/handler"
	"github.com/Kei-K23/go-url-shortened/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB


func init() {
	var err error
	db, err = gorm.Open(sqlite.Open("urls.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	db.AutoMigrate(&models.Url{})
	fmt.Println("Successfully connected to database")
	fmt.Println("Server is running on http://localhost:8080")
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /urls", func(w http.ResponseWriter, r *http.Request) {
		handler.GetURLs(w, r, db)
	})
	
	mux.HandleFunc("POST /urls/redirect/", func(w http.ResponseWriter, r *http.Request) {
		handler.RedirectTo(w, r, db)
	})

	mux.HandleFunc("POST /urls", func(w http.ResponseWriter, r *http.Request) {
		handler.CreateShortURL(w, r, db)
	})

	log.Fatal(http.ListenAndServe(":8080", mux))
}
