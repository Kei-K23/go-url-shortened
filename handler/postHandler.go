package handler

import (
	"encoding/json"
	"fmt"
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

type CreateBodyPayload struct {
	OriginalUrl string `json:"original_url"`
}

type RedirectBodyPayload struct {
	ShortenedUrl string `json:"shortened_url"`
}

func CreateShortURL(w http.ResponseWriter, r *http.Request, db *gorm.DB) {

	var newUrl string

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	var data CreateBodyPayload

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


func RedirectTo(w http.ResponseWriter, r *http.Request , db *gorm.DB) {
	fmt.Println("hit")
	var url models.Url

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	var data RedirectBodyPayload

	if err := json.Unmarshal(body, &data); err != nil {
		http.Error(w, "Error when parsing json", http.StatusBadRequest)
	}

	shortUrl := data.ShortenedUrl


	result := db.Where("shortened_url = ?" , shortUrl).First(&url)

	if result.Error!= nil {
        http.Error(w, result.Error.Error(), http.StatusNotFound)
        return
    }

	http.Redirect(w, r, url.OriginalUrl, http.StatusFound)
}