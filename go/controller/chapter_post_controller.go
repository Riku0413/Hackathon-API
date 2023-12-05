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

func ChapterPostHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")

	switch r.Method {
	case http.MethodPost:
		decoder := json.NewDecoder(r.Body)
		var newChapter model.Chapter
		if err := decoder.Decode(&newChapter); err != nil {
			log.Printf("fail: json decode, %v\n", err)
			log.Printf("fail: controller.ChapterPostHandler(), %v\n", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		ms := ulid.Timestamp(time.Now())
		entropy := rand.New(rand.NewSource(time.Now().UnixNano()))
		ulidValue, err := ulid.New(ms, entropy)
		if err != nil {
			log.Printf("fail: ulid generate, %v\n", err)
			log.Printf("fail: controller.ChapterPostHandler(), %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		newChapter.Id = ulidValue.String()

		if err := newChapter.ChapterCheckUID(); err != nil {
			log.Printf("fail: chapter.ChpaterCheckUID(), %v\n", err)
			log.Printf("fail: controller.ChapterPostHandler(), %v\n", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		now := time.Now()
		seconds := now.Unix()
		t := time.Unix(seconds, 0)
		newChapter.BirthTime = t.Format("2006-01-02 15:04:05")
		newChapter.UpdateTime = t.Format("2006-01-02 15:04:05")

		if err := usecase.ChapterPost(newChapter); err != nil {
			log.Printf("fail: usecase.ChapterPost, %v\n", err)
			log.Printf("fail: controller.ChapterPostHandler(), %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		response := map[string]interface{}{
			"id": newChapter.Id,
		}
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			log.Printf("fail: json.Marshal, %v\n", err)
			log.Printf("fail: controller.ChapterPostHandler(), %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
		log.Printf("post: controller.ChapterPostHandler(), %v\n", err)

	default:
		log.Printf("fail: controller.ChapterPostHandler(), HTTP Method is %s\n", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
