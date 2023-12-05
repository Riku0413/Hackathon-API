package controller

import (
	"encoding/json"
	"github.com/Riku0413/Hackathon-API/usecase"
	"log"
	"net/http"
)

func BooksSearchHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")

	switch r.Method {
	case http.MethodGet:
		queryParams := r.URL.Query()
		keyword := queryParams.Get("q")
		if keyword == "" {
			log.Printf("fail: controller.BooksSearchHandler(), keyword query is null")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		books, err := usecase.BooksSearch(keyword)
		if err != nil {
			log.Printf("fail: usecase.BooksSearch(), %v\n", err)
			log.Printf("fail: controller.BooksSearchHandler(), %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		jsonResponses, err := json.Marshal(books)
		if err != nil {
			log.Printf("fail: json.Marshal(), %v\n", err)
			log.Printf("fail: controller.BooksSearchHandler(), %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponses)

	default:
		log.Printf("fail: controller.BooksSearchHandler(), HTTP Method is %s\n", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
