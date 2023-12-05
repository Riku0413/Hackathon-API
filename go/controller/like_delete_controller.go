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

func LikeDeleteHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")

	switch r.Method {
	case http.MethodDelete:
		parsedURL, err := url.Parse(r.URL.String())
		if err != nil {
			log.Printf("fail: controller.LikeDeleteHandler(), %v\n", err)
			return
		}
		var like model.Like
		segments := strings.Split(parsedURL.Path, "/")
		like.ItemCategory = segments[len(segments)-3]
		like.ItemId = segments[len(segments)-2]
		like.Id = segments[len(segments)-1]

		if err := like.LikeCheckUID(); err != nil {
			log.Printf("fail: model.Like.LikeCheckUID(), %v\n", err)
			log.Printf("fail: controller.LikeDeleteHandler(), %v\n", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err := usecase.LikeDelete(like.Id, like.ItemCategory, like.ItemId); err != nil {
			log.Printf("fail: usecase.LikeDelete(), %v\n", err)
			log.Printf("fail: controller.LikeDeleteHandler(), %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		response := map[string]interface{}{
			"id": like.Id,
		}
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			log.Printf("fail: json.Marshal(), %v\n", err)
			log.Printf("fail: controller.LikeDeleteHandler(), %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// フロントにレスポンスとして登録したIDを返す
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
		log.Printf("delete: controller.LikeDeleteHandler(), %v\n", err)

	default:
		log.Printf("fail: controller.LikeDeleteHandler(), HTTP Method is %s\n", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
