package controller

import (
	"encoding/json"
	_ "fmt"
	"github.com/Riku0413/Hackathon-API/model"
	"github.com/Riku0413/Hackathon-API/usecase"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func BookDeleteHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")

	switch r.Method {
	case http.MethodDelete:
		parsedURL, err := url.Parse(r.URL.String())
		if err != nil {
			log.Printf("fail: controller.BookDeleteHandler(), %v\n", err)
			return
		}
		var book model.Book
		segments := strings.Split(parsedURL.Path, "/")
		book.Id = segments[len(segments)-1]

		if err := book.BookCheckUID(); err != nil {
			log.Printf("fail: model.Book.BookCheckUID(), %v\n", err)
			log.Printf("fail: controller.BookDeleteHandler(), %v\n", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err := usecase.BookDelete(book.Id); err != nil {
			log.Printf("fail: usecase.BookDelete(), %v\n", err)
			log.Printf("fail: controller.BookDeleteHandler(), %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		response := map[string]interface{}{
			"id": book.Id,
		}
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			log.Printf("fail: json.Marshal(), %v\n", err)
			log.Printf("fail: controller.BookDeleteHandler(), %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// フロントにレスポンスとして登録したIDを返す
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
		log.Printf("delete: controller.BookDeleteHandler(), %v\n", err)

	default:
		log.Printf("fail: controller.BookDeleteHandler(), HTTP Method is %s\n", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
