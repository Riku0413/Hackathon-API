package controller

import (
	"log"
	"net/http"
)

// アクセス権の付与！
func OptionsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*") // ここの条件は審議！ 
	w.Header().Set("Access-Control-Allow-Headers", "*") // ここのスペルミス！！！
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

	switch r.Method {
	case http.MethodOptions:
		log.Printf("options 1") // これいる？　いるなら統一しよう
		w.WriteHeader(http.StatusOK)
		return
	default:
		log.Printf("fail: HTTP Method is %s\n", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}