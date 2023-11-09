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

func VideoGetHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")

	switch r.Method {
	case http.MethodGet:
		parsedURL, err := url.Parse(r.URL.String())
		if err != nil {
			log.Printf("fail: url.Parse(), %v\n", err)
			log.Printf("fail: controller.VideoGetHandler(), %v\n", err)
			return
		}
		var video model.Video
		segments := strings.Split(parsedURL.Path, "/")
		video.Id = segments[len(segments)-1]

		if err := video.VideoCheckUID(); err != nil {
			log.Printf("fail: model.Video.VideoCheckUID(), %v\n", err)
			log.Printf("fail: controller.VideoGetHandler(), %v\n", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		video, err = usecase.VideoGet(video.Id)
		if err != nil {
			log.Printf("fail: usecase.VideoGet(), %v\n", err)
			log.Printf("fail: controller.VideoGetHandler(), %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		jsonResponse, err := json.Marshal(video)
		if err != nil {
			log.Printf("fail: json.Marshal(), %v\n", err)
			log.Printf("fail: controller.VideoGetHandler(), %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
		log.Printf("get: controller.VideoGetHandler()")

	default:
		log.Printf("fail: controller.VideoGetHandler(), HTTP Method is %s\n", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
