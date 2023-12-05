package controller

import (
	"encoding/json"
	_ "fmt"
	"github.com/Riku0413/Hackathon-API/model"
	"github.com/Riku0413/Hackathon-API/usecase"
	"log"
	"net/http"
	"net/url"
	"time"
)

func VideoPutHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")

	switch r.Method {
	case http.MethodPut:
		decoder := json.NewDecoder(r.Body)
		var jsonInput map[string]string
		if err := decoder.Decode(&jsonInput); err != nil {
			log.Printf("fail: json decode, %v\n", err)
			log.Printf("fail: controller.VideoPutHandler(), %v\n", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		// URL フィールドの値を URL 型に変換
		urlString, ok := jsonInput["url"]
		if !ok {
			log.Printf("fail: controller.VideoPutHandler(), URL not found in JSON input\n")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		videoLink, err := url.Parse(urlString)
		if err != nil {
			log.Printf("fail: URL is invalid, %v\n", err)
			log.Printf("fail: controller.VideoPutHandler(), %v\n", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		var video model.Video
		video.Id = jsonInput["id"]
		video.Title = jsonInput["title"]
		video.Introduction = jsonInput["introduction"]
		video.URL = videoLink

		if err := video.VideoCheckUID(); err != nil {
			log.Printf("fail: model.Video.VideoCheckUID, %v\n", err)
			log.Printf("fail: controller.VideoPutHandler(), %v\n", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		now := time.Now()
		seconds := now.Unix()
		t := time.Unix(seconds, 0)
		video.UpdateTime = t.Format("2006-01-02 15:04:05")

		if err := usecase.VideoPut(video); err != nil {
			log.Printf("fail: usecase.VideoPut(), %v\n", err)
			log.Printf("fail: controller.VideoPutHandler(), %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		response := map[string]interface{}{
			"id": video.Id,
		}
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			log.Printf("fail: json.Marshal, %v\n", err)
			log.Printf("fail: controller.VideoPutHandler(), %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
		log.Printf("post: controller.VideoPutHandler(), %v\n", err)

	default:
		log.Printf("fail: controller.VideoPutHandler(), HTTP Method is %s\n", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
