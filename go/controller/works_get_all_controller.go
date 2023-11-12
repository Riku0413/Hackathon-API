package controller

import (
	"encoding/json"
	"github.com/Riku0413/Hackathon-API/usecase"
	"log"
	"net/http"
)

func WorksGetAllHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")

	switch r.Method {
	case http.MethodGet:
		works, err := usecase.WorksGetAll()
		if err != nil {
			log.Printf("fail: usecase.WorksGetAll(), %v\n", err)
			log.Printf("fail: controller.WorksGetAllHandler(), %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		jsonResponses, err := json.Marshal(works)
		if err != nil {
			log.Printf("fail: json.Marshal(), %v\n", err)
			log.Printf("fail: controller.WorksGetAllHandler(), %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponses)
		log.Printf("get: controller.WorksGetAllHandler()")

	default:
		log.Printf("fail: controller.WorksGetAllHandler(), HTTP Method is %s\n", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
