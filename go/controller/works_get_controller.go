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

func WorksGetHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")

	switch r.Method {
	case http.MethodGet:
		parsedURL, err := url.Parse(r.URL.String())
		if err != nil {
			log.Printf("fail: url.Parse(), %v\n", err)
			log.Printf("fail: controller.WorksGetHandler(), %v\n", err)
			return
		}
		var user model.User
		segments := strings.Split(parsedURL.Path, "/")
		user.Id = segments[len(segments)-1]

		if err := user.UserCheckUID(); err != nil {
			log.Printf("fail: model.User.UserCheckUID(), %v\n", err)
			log.Printf("fail: controller.WorksGetHandler(), %v\n", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		works, err := usecase.WorksGet(user.Id)
		if err != nil {
			log.Printf("fail: usecase.WorksGet(), %v\n", err)
			log.Printf("fail: controller.WorksGetHandler(), %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		jsonResponses, err := json.Marshal(works)
		if err != nil {
			log.Printf("fail: json.Marshal(), %v\n", err)
			log.Printf("fail: controller.WorksGetHandler(), %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponses)
		log.Printf("get: controller.WorksGetHandler()")

	default:
		log.Printf("fail: controller.WorksGetHandler(), HTTP Method is %s\n", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
