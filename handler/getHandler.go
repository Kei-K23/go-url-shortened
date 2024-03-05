package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Kei-K23/go-url-shortened/models"
	"gorm.io/gorm"
)

type Response struct {
	ID        uint `json:"id"`
	OriginalUrl  string `json:"original_url"`
	ShortenedUrl string `json:"shortened_url"`
		CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}


func GetURLs(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	
	var urls []models.Url
	var resData []Response
	result := db.Find(&urls)

	if result.Error!= nil {
        http.Error(w, result.Error.Error(), http.StatusInternalServerError)
        return
    }

	if len(urls) == 0 {
		http.Error(w, "No URLs found", http.StatusNotFound)
        return
	}

	for _, url := range urls  { 
		resData = append(resData, Response{
			ID: url.ID,
			OriginalUrl: url.OriginalUrl,
		ShortenedUrl: url.ShortenedUrl,
	CreatedAt: url.CreatedAt,
	UpdatedAt: url.UpdatedAt,	
	})
	}

	resJsonPayload , err := json.Marshal(resData)

	if err!= nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resJsonPayload)
}

