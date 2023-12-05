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

func BlogPutHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")

	switch r.Method {
	case http.MethodPut:
		decoder := json.NewDecoder(r.Body)
		var blog model.Blog
		if err := decoder.Decode(&blog); err != nil {
			log.Printf("fail: json decode, %v\n", err)
			log.Printf("fail: controller.BlogPutHandler(), %v\n", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		if err := blog.BlogCheckUID(); err != nil {
			log.Printf("fail: model.Blog.BlogCheckUID(), %v\n", err)
			log.Printf("fail: controller.BlogPutHandler(), %v\n", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		now := time.Now()
		seconds := now.Unix()
		t := time.Unix(seconds, 0)
		blog.UpdateTime = t.Format("2006-01-02 15:04:05")

		if err := usecase.BlogPut(blog); err != nil {
			log.Printf("fail: usecase.BlogPut(), %v\n", err)
			log.Printf("fail: controller.BlogPutHandler(), %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		response := map[string]interface{}{
			"id": blog.Id,
		}
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			log.Printf("fail: json.Marshal(), %v\n", err)
			log.Printf("fail: controller.BlogPutHandler(), %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
		log.Printf("put: controller.BlogPutHandler()")

	default:
		log.Printf("fail: controller.BlogPutHandler(), HTTP Method is %s\n", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
