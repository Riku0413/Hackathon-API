package controller

import (
	"encoding/json"
	_ "fmt"
	"github.com/Riku0413/Hackathon-API/model"
	"github.com/Riku0413/Hackathon-API/usecase"
	"github.com/oklog/ulid/v2"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func CommentPostHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")

	switch r.Method {
	case http.MethodPost:
		decoder := json.NewDecoder(r.Body)
		var newComment model.Comment
		if err := decoder.Decode(&newComment); err != nil {
			log.Printf("fail: json decode, %v\n", err)
			log.Printf("fail: controller.CommentPostHandler(), %v\n", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		ms := ulid.Timestamp(time.Now())
		entropy := rand.New(rand.NewSource(time.Now().UnixNano()))
		ulidValue, err := ulid.New(ms, entropy)
		if err != nil {
			log.Printf("fail: ulid generate, %v\n", err)
			log.Printf("fail: controller.CommentPostHandler(), %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		newComment.Id = ulidValue.String()

		if err := newComment.CommentCheckUID(); err != nil {
			log.Printf("fail: model.Comment.CommentCheckUID(), %v\n", err)
			log.Printf("fail: controller.CommentPostHandler(), %v\n", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		now := time.Now()
		seconds := now.Unix()
		t := time.Unix(seconds, 0)
		newComment.BirthTime = t.Format("2006-01-02 15:04:05")
		newComment.UpdateTime = t.Format("2006-01-02 15:04:05")

		if err := usecase.CommentPost(newComment); err != nil {
			log.Printf("fail: usecase.CommentPost(), %v\n", err)
			log.Printf("fail: controller.CommentPostHandler(), %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		response := map[string]interface{}{
			"id": newComment.Id,
		}
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			log.Printf("fail: json.Marshal(), %v\n", err)
			log.Printf("fail: controller.CommentPostHandler(), %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
		log.Printf("post: controller.CommentPostHandler()")

	default:
		log.Printf("fail: controller.CommentPostHandler(), HTTP Method is %s\n", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
