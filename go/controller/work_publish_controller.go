package controller

import (
	"encoding/json"
	_ "fmt"
	"github.com/Riku0413/Hackathon-API/model"
	"github.com/Riku0413/Hackathon-API/usecase"
	"log"
	"net/http"
	"time"
)

func WorkPublishHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")

	switch r.Method {
	case http.MethodPut:
		decoder := json.NewDecoder(r.Body)
		var work model.Work
		if err := decoder.Decode(&work); err != nil {
			log.Printf("fail: json decode, %v\n", err)
			log.Printf("fail: controller.WorkPublishHandler(), %v\n", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		if err := work.WorkCheckUID(); err != nil {
			log.Printf("fail: model.Work.WorkCheckUID(), %v\n", err)
			log.Printf("fail: controller.WorkPublishHandler(), %v\n", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		now := time.Now()
		seconds := now.Unix()
		t := time.Unix(seconds, 0)
		work.UpdateTime = t.Format("2006-01-02 15:04:05")

		if err := usecase.WorkPublish(work); err != nil {
			log.Printf("fail: usecase.WorkPublish(), %v\n", err)
			log.Printf("fail: controller.WorkPublishHandler(), %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		response := map[string]interface{}{
			"id": work.Id,
		}
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			log.Printf("fail: json.Marshal(), %v\n", err)
			log.Printf("fail: controller.WorkPublishHandler(), %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
		log.Printf("post: controller.WorkPublishHandler(), %v\n", err)

	default:
		log.Printf("fail: controller.WorkPublishHandler(), HTTP Method is %s\n", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
