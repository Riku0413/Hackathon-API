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

func WorkGetHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")

	switch r.Method {
	case http.MethodGet:
		parsedURL, err := url.Parse(r.URL.String())
		if err != nil {
			log.Printf("fail: url.Parse(), %v\n", err)
			log.Printf("fail: controller.WorkGetHandler(), %v\n", err)
			return
		}
		var work model.Work
		segments := strings.Split(parsedURL.Path, "/")
		work.Id = segments[len(segments)-1]

		if err := work.WorkCheckUID(); err != nil {
			log.Printf("fail: model.Work.WorkCheckUID(), %v\n", err)
			log.Printf("fail: controller.WorkGetHandler(), %v\n", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		work, err = usecase.WorkGet(work.Id)
		if err != nil {
			log.Printf("fail: usecase.WorkGet(), %v\n", err)
			log.Printf("fail: controller.WorkGetHandler(), %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		jsonResponse, err := json.Marshal(work)
		if err != nil {
			log.Printf("fail: json.Marshal(), %v\n", err)
			log.Printf("fail: controller.WorkGetHandler(), %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
		log.Printf("get: controller.WorkGetHandler()")

	default:
		log.Printf("fail: controller.WorkGetHandler(), HTTP Method is %s\n", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
