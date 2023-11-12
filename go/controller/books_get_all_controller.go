package controller

import (
	"encoding/json"
	"github.com/Riku0413/Hackathon-API/usecase"
	"log"
	"net/http"
)

func BooksGetAllHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")

	switch r.Method {
	case http.MethodGet:
		books, err := usecase.BooksGetAll()
		if err != nil {
			log.Printf("fail: usecase.BooksGetAll(), %v\n", err)
			log.Printf("fail: controller.BooksGetAllHandler(), %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		jsonResponses, err := json.Marshal(books)
		if err != nil {
			log.Printf("fail: json.Marshal(), %v\n", err)
			log.Printf("fail: controller.BooksGetAllHandler(), %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponses)
		log.Printf("get: controller.BooksGetAllHandler()")

	default:
		log.Printf("fail: controller.BooksGetAllHandler(), HTTP Method is %s\n", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
