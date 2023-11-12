package controller

import (
	"encoding/json"
	_ "fmt"
	"github.com/Riku0413/Hackathon-API/model"
	"github.com/Riku0413/Hackathon-API/usecase"
	"log"
	"net/http"
)

func UserUpdateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")

	switch r.Method {
	case http.MethodPut:
		decoder := json.NewDecoder(r.Body)
		var user model.User
		if err := decoder.Decode(&user); err != nil {
			log.Printf("fail: json decode, %v\n", err)
			log.Printf("fail: controller.UserUpdateHandler(), %v\n", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		if err := user.UserCheckUID(); err != nil {
			log.Printf("fail: model.User.UserCheckUID(), %v\n", err)
			log.Printf("fail: controller.UserUpdateHandler(), %v\n", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err := usecase.UserUpdate(user); err != nil {
			log.Printf("fail: usecase.UserUpdate(), %v\n", err)
			log.Printf("fail: controller.UserUpdateHandler(), %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		response := map[string]interface{}{
			"id": user.Id,
		}
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			log.Printf("fail: json.Marshal(), %v\n", err)
			log.Printf("fail: controller.UserUpdateHandler(), %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
		log.Printf("put: controller.UserUpdateHandler()")

	default:
		log.Printf("fail: controller.UserUpdateHandler(), HTTP Method is %s\n", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
