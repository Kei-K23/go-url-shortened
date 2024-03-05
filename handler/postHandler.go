package handler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/Kei-K23/go-url-shortened/lib"
	"github.com/Kei-K23/go-url-shortened/models"
	"gorm.io/gorm"
)

type CreateRes struct {
	OriginalUrl  string `json:"original_url"`
	ShortenedUrl string `json:"shortened_url"`
}

type BodyPayload struct {
	OriginalUrl string `json:"original_url"`
}

func CreateShortURL(w http.ResponseWriter, r *http.Request, db *gorm.DB) {

	var newUrl string

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	var data BodyPayload

	if err := json.Unmarshal(body, &data); err != nil {
		http.Error(w, "Error when parsing json", http.StatusBadRequest)
	}

	originalUrl := data.OriginalUrl

	parsedUrl, err := url.Parse(originalUrl)
	if err != nil {
		log.Fatal("Error when parsing URL : ", originalUrl)
	}

	if parsedUrl.Scheme == "https" {
		newUrl = "https://" + lib.GenerateRandomString(10)
	} else {
		newUrl = "http://" + lib.GenerateRandomString(10)
	}

	payload := CreateRes{
		OriginalUrl:  originalUrl,
		ShortenedUrl: newUrl,
	}

	w.Header().Set("Content-Type", "application/json")

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	url := &models.Url{
        OriginalUrl: originalUrl,
        ShortenedUrl: newUrl,
    }

	result := db.Create(url)
	if result.Error!= nil {
        log.Fatal(result.Error)
    }

	w.Write(jsonPayload)	
}
