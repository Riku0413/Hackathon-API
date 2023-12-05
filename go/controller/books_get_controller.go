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

func BooksGetHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")

	switch r.Method {
	case http.MethodGet:
		parsedURL, err := url.Parse(r.URL.String())
		if err != nil {
			log.Printf("fail: url.Parse(), %v\n", err)
			log.Printf("fail: controller.BooksGetHandler(), %v\n", err)
			return
		}
		var user model.User
		segments := strings.Split(parsedURL.Path, "/")
		user.Id = segments[len(segments)-1]

		if err := user.UserCheckUID(); err != nil {
			log.Printf("fail: model.User.UserCheckUID(), %v\n", err)
			log.Printf("fail: controller.BooksGetHandler(), %v\n", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		books, err := usecase.BooksGet(user.Id)
		if err != nil {
			log.Printf("fail: usecase.BooksGet(), %v\n", err)
			log.Printf("fail: controller.BooksGetHandler(), %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		jsonResponses, err := json.Marshal(books)
		if err != nil {
			log.Printf("fail: json.Marshal(), %v\n", err)
			log.Printf("fail: controller.BooksGetHandler(), %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponses)
		log.Printf("get: controller.BooksGetHandler()")

	default:
		log.Printf("fail: controller.BooksGetHandler(), HTTP Method is %s\n", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
