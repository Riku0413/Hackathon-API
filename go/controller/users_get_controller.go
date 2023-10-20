package controller

import (
	"encoding/json"
	"github.com/Riku0413/Hackathon-API/usecase"
	"log"
	"net/http"
)

func GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	// アクセス権を追加
	// ここいるのかな
	w.Header().Set("Access-Control-Allow-Origin", "*") // ここの条件は審議！
	w.Header().Set("Access-Control-Allow-Headers", "*") // ここのスペルミス！！！
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

	switch r.Method {
	// これがいるかどうかは後で動作確認してチェック
	// case http.MethodOptions:
	// 	log.Printf("options 3") // これいる？　いるなら統一しよう
	// 	w.WriteHeader(http.StatusOK)
	// 	return

	case http.MethodGet:
		log.Printf("get") // これいる？　いるなら統一しよう
		users, err := usecase.GetAllUsers()
		if err != nil {
			log.Printf("fail: usecase.GetAllUsers, %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		bytes, err := json.Marshal(users)
		if err != nil {
			log.Printf("fail: json.Marshal, %v\n", err)
			w.WriteHeader(http.StatusInternalServerError) // これあってる？？
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(bytes) // ここら辺のエラーは？？

	default:
		log.Printf("fail: HTTP Method is %s\n", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
