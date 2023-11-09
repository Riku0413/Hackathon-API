package controller

import (
	"encoding/json"
	_ "fmt"
	"github.com/Riku0413/Hackathon-API/model"
	"github.com/Riku0413/Hackathon-API/usecase"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func BlogDeleteHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")

	switch r.Method {
	case http.MethodDelete:
		parsedURL, err := url.Parse(r.URL.String())
		if err != nil {
			log.Printf("fail: url.parse(), %v\n", err)
			log.Printf("fail: controller.BlogDeleteHandler(), %v\n", err)
			return
		}
		var blog model.Blog
		segments := strings.Split(parsedURL.Path, "/")
		blog.Id = segments[len(segments)-1]

		if err := blog.BlogCheckUID(); err != nil {
			log.Printf("fail: model.Blog.BlogCheckUID(), %v\n", err)
			log.Printf("fail: controller.BlogDeleteHandler(), %v\n", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err := usecase.BlogDelete(blog.Id); err != nil {
			log.Printf("fail: usecase.BlogDelete(), %v\n", err)
			log.Printf("fail: controller.BlogDeleteHandler(), %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// クライアントへのレスポンスを作成し、JSON形式に変換
		// JSON形式はcontrollerでしか登場させない！
		response := map[string]interface{}{
			"id": blog.Id, // ここのエラーはないの？
		}
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			log.Printf("fail: json.Marshal(), %v\n", err)
			log.Printf("fail: controller.BlogDeleteHandler(), %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// フロントにレスポンスとして登録したIDを返す
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
		log.Printf("delete: controller.BlogDeleteHandler()")

	default:
		log.Printf("fail: controller.BlogDeleteHandler(), HTTP Method is %s\n", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
