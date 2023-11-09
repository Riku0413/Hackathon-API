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

func ChaptersGetHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")

	switch r.Method {
	case http.MethodGet:
		parsedURL, err := url.Parse(r.URL.String())
		if err != nil {
			log.Printf("fail: url.Parse(), %v\n", err)
			log.Printf("fail: controller.ChapterGetHandler(), %v\n", err)
			return
		}
		var book model.Book
		segments := strings.Split(parsedURL.Path, "/")
		book.Id = segments[len(segments)-1]

		if err := book.BookCheckUID(); err != nil {
			log.Printf("fail: model.Book.BookCheckUID(), %v\n", err)
			log.Printf("fail: controller.ChaptersGetHandler(), %v\n", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		chapters, err := usecase.ChaptersGet(book.Id)
		if err != nil {
			log.Printf("fail: usecase.ChaptersGet(), %v\n", err)
			log.Printf("fail: controller.ChaptersGetHandler(), %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		jsonResponses, err := json.Marshal(chapters)
		if err != nil {
			log.Printf("fail: json.Marshal, %v\n", err)
			log.Printf("fail: controller.ChaptersGetHandler(), %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponses)
		log.Printf("get: controller.ChaptersGetHandler()")

	default:
		log.Printf("fail: controller.ChaptersGetHandler(), HTTP Method is %s\n", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
