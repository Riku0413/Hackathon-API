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

func BlogPostHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")

	switch r.Method {
	case http.MethodPost:
		decoder := json.NewDecoder(r.Body)
		var newBlog model.Blog
		if err := decoder.Decode(&newBlog); err != nil {
			log.Printf("fail: json decode, %v\n", err)
			log.Printf("fail: controller.BlogPostHandler(), %v\n", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		ms := ulid.Timestamp(time.Now())
		entropy := rand.New(rand.NewSource(time.Now().UnixNano()))
		ulidValue, err := ulid.New(ms, entropy)
		if err != nil {
			log.Printf("fail: ulid generate, %v\n", err)
			log.Printf("fail: controller.BlogPostHandler(), %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		newBlog.Id = ulidValue.String()

		if err := newBlog.BlogCheckUID(); err != nil {
			log.Printf("fail: model.Blog.BlogCheckUID(), %v\n", err)
			log.Printf("fail: controller.BlogPostHandler(), %v\n", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		now := time.Now()
		seconds := now.Unix()
		t := time.Unix(seconds, 0)
		newBlog.BirthTime = t.Format("2006-01-02 15:04:05")
		newBlog.UpdateTime = t.Format("2006-01-02 15:04:05")

		if err := usecase.BlogPost(newBlog); err != nil {
			log.Printf("fail: usecase.BlogPost(), %v\n", err)
			log.Printf("fail: controller.BlogPostHandler(), %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		response := map[string]interface{}{
			"id": newBlog.Id,
		}
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			log.Printf("fail: json.Marshal(), %v\n", err)
			log.Printf("fail: controller.BlogPostHandler(), %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
		log.Printf("post: controller.BlogPostHandler()")

	default:
		log.Printf("fail: controller.BlogPostHandler(), HTTP Method is %s\n", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
