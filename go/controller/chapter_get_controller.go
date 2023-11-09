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

func ChapterGetHandler(w http.ResponseWriter, r *http.Request) {
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
		var chapter model.Chapter
		segments := strings.Split(parsedURL.Path, "/")
		chapter.Id = segments[len(segments)-1]

		if err := chapter.ChapterCheckUID(); err != nil {
			log.Printf("fail: model.Chapter.ChapterCheckUID(), %v\n", err)
			log.Printf("fail: controller.ChapterGetHandler(), %v\n", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		chapter, err = usecase.ChapterGet(chapter.Id)
		if err != nil {
			log.Printf("fail: usecase.ChapterGet(), %v\n", err)
			log.Printf("fail: controller.ChapterGetHandler(), %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		jsonResponse, err := json.Marshal(chapter)
		if err != nil {
			log.Printf("fail: json.Marshal(), %v\n", err)
			log.Printf("fail: controller.ChapterGetHandler(), %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
		log.Printf("get: controller.ChapterGetHandler()")

	default:
		log.Printf("fail: controller.ChapterGetHandler(), HTTP Method is %s\n", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
