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

func BlogGetHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")

	switch r.Method {
	case http.MethodGet:
		parsedURL, err := url.Parse(r.URL.String())
		if err != nil {
			log.Printf("fail: url.Parse(), %v\n", err)
			log.Printf("fail: controller.BlogGetHandler(), %v\n", err)
			return
		}
		var blog model.Blog
		segments := strings.Split(parsedURL.Path, "/")
		blog.Id = segments[len(segments)-1]

		if err := blog.BlogCheckUID(); err != nil {
			log.Printf("fail: model.Blog.BlogCheckUID(), %v\n", err)
			log.Printf("fail: controller.BlogGetHandler(), %v\n", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		blog, err = usecase.BlogGet(blog.Id)
		if err != nil {
			log.Printf("fail: usecase.BlogGet(), %v\n", err)
			log.Printf("fail: controller.BlogGetHandler(), %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		jsonResponse, err := json.Marshal(blog)
		if err != nil {
			log.Printf("fail: json.Marshal(), %v\n", err)
			log.Printf("fail: controller.BlogGetHandler(), %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
		log.Printf("get: controller.BlogGetHandler()")

	default:
		log.Printf("fail: controller.BlogGetHandler(), HTTP Method is %s\n", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
