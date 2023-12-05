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

func VideoDeleteHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")

	switch r.Method {
	case http.MethodDelete:
		parsedURL, err := url.Parse(r.URL.String())
		if err != nil {
			log.Printf("fail: url.parse(), %v\n", err)
			log.Printf("fail: controller.VideoDeleteHandler(), %v\n", err)
			return
		}
		var video model.Video
		segments := strings.Split(parsedURL.Path, "/")
		video.Id = segments[len(segments)-1]

		if err := video.VideoCheckUID(); err != nil {
			log.Printf("fail: model.Video.VideoCheckUID(), %v\n", err)
			log.Printf("fail: controller.VideoDeleteHandler(), %v\n", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err := usecase.VideoDelete(video.Id); err != nil {
			log.Printf("fail: usecase.VideoDelete(), %v\n", err)
			log.Printf("fail: controller.VideoDeleteHandler(), %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		response := map[string]interface{}{
			"id": video.Id,
		}
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			log.Printf("fail: json.Marshal(), %v\n", err)
			log.Printf("fail: controller.VideoDeleteHandler(), %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
		log.Printf("delete: controller.VideoDeleteHandler()")

	default:
		log.Printf("fail: controller.VideoDeleteHandler(), HTTP Method is %s\n", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
