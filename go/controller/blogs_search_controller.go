package controller

import (
	"encoding/json"
	"github.com/Riku0413/Hackathon-API/usecase"
	"log"
	"net/http"
)

func BlogsSearchHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")

	switch r.Method {
	case http.MethodGet:
		queryParams := r.URL.Query()
		keyword := queryParams.Get("q")
		if keyword == "" {
			log.Printf("fail: controller.BlogsSearchHandler(), keyword query is null")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		blogs, err := usecase.BlogsSearch(keyword)
		if err != nil {
			log.Printf("fail: usecase.BlogsSearch(), %v\n", err)
			log.Printf("fail: controller.BlogsearchHandler(), %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		jsonResponses, err := json.Marshal(blogs)
		if err != nil {
			log.Printf("fail: json.Marshal(), %v\n", err)
			log.Printf("fail: controller.BlogsearchHandler(), %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponses)

	default:
		log.Printf("fail: controller.BlogsearchHandler(), HTTP Method is %s\n", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
