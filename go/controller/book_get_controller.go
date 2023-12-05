package controller

import (
	"encoding/json"
	"github.com/Riku0413/Hackathon-API/model"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/Riku0413/Hackathon-API/usecase"
)

func BookGetHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")

	switch r.Method {
	case http.MethodGet:
		parsedURL, err := url.Parse(r.URL.String())
		if err != nil {
			log.Printf("fail: url.Parse(), %v\n", err)
			log.Printf("fail: controller.BookGetHandler(), %v\n", err)
			return
		}
		var book model.Book
		segments := strings.Split(parsedURL.Path, "/")
		book.Id = segments[len(segments)-1]

		if err := book.BookCheckUID(); err != nil {
			log.Printf("fail: model.Book.BookCheckUID(), %v\n", err)
			log.Printf("fail: controller.BookGetHandler(), %v\n", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		book, err = usecase.BookGet(book.Id)
		if err != nil {
			log.Printf("fail: usecase.BookGet(), %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		jsonResponse, err := json.Marshal(book)
		if err != nil {
			log.Printf("fail: json.Marshal(), %v\n", err)
			log.Printf("fail: controller.BookGetHandler(), %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
		log.Printf("get: controller.BookGetHandler()")

	default:
		log.Printf("fail: controller.BookGetHandler(), HTTP Method is %s\n", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
