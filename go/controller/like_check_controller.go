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

func LikeCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")

	switch r.Method {
	case http.MethodGet:
		parsedURL, err := url.Parse(r.URL.String())
		if err != nil {
			log.Printf("fail: url.Parse(), %v\n", err)
			log.Printf("fail: controller.LikeCheckHandler(), %v\n", err)
			return
		}
		var like model.Like
		segments := strings.Split(parsedURL.Path, "/")
		like.ItemCategory = segments[len(segments)-3]
		like.ItemId = segments[len(segments)-2]
		like.UserId = segments[len(segments)-1]

		like, err = usecase.LikeCheck(like.ItemCategory, like.ItemId, like.UserId)
		if err != nil {
			log.Printf("fail: usecase.LikeCheck(), %v\n", err)
			log.Printf("fail: controller.LikeCheckHandler(), %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		jsonResponse, err := json.Marshal(like)
		if err != nil {
			log.Printf("fail: json.Marshal(), %v\n", err)
			log.Printf("fail: controller.LikeCheckHandler(), %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
		log.Printf("get: controller.LikeCheckHandler()")

	default:
		log.Printf("fail: controller.LikeCheckHandler(), HTTP Method is %s\n", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
