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

func CommentsGetHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")

	switch r.Method {
	case http.MethodGet:
		parsedURL, err := url.Parse(r.URL.String())
		if err != nil {
			log.Printf("fail: url.Parse(), %v\n", err)
			log.Printf("fail: controller.CommentsGetHandler(), %v\n", err)
			return
		}
		var comment model.Comment
		segments := strings.Split(parsedURL.Path, "/")
		comment.ItemCategory = segments[len(segments)-2]
		comment.ItemId = segments[len(segments)-1]

		//if err := comment.BlogCheckUID(); err != nil {
		//	log.Printf("fail: model.Comment.CommentCheckUID(), %v\n", err)
		//	log.Printf("fail: controller.CommentsGetHandler(), %v\n", err)
		//	w.WriteHeader(http.StatusBadRequest)
		//	return
		//}

		comments, err := usecase.CommentsGet(comment.ItemCategory, comment.ItemId)
		if err != nil {
			log.Printf("fail: usecase.CommentsGet(), %v\n", err)
			log.Printf("fail: controller.CommentsGetHandler(), %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		jsonResponses, err := json.Marshal(comments)
		if err != nil {
			log.Printf("fail: json.Marshal(), %v\n", err)
			log.Printf("fail: controller.CommentsGetHandler(), %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponses)
		log.Printf("get: controller.CommentsGetHandler()")

	default:
		log.Printf("fail: controller.CommentsGetHandler(), HTTP Method is %s\n", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
