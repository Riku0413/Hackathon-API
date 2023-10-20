package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Riku0413/Hackathon-API/usecase"
)

// ② nameパラメータと一致するデータを返す & 新データのポスト
func SearchUserHandler(w http.ResponseWriter, r *http.Request) {
	// ここいるのかな
	w.Header().Set("Access-Control-Allow-Origin", "*") // ここの条件は審議！
	w.Header().Set("Access-Control-Allow-Headers", "*") // ここのスペルミス！！！
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	
	switch r.Method {
	case http.MethodGet:
		// ②-1　リクエストのクエリパラメータを取得
		queryParams := r.URL.Query()
		name := queryParams.Get("name") // ここでエラーにならないの？？
		// バリデーションはcontrollerで行う！
		if name == "" {
			log.Println("fail: name is empty")
			w.WriteHeader(http.StatusBadRequest) // httpリクエストに関するエラーはここ（controller）で弾く！
			return
		}

		// データをGo特有の型に整形し、バリデーションをかけるまでがcontrollerの仕事！
		// 完了したら、正当なデータをusecaseに送る、そして結果を受け取る
		users, err := usecase.GetUsersByName(name)
		if err != nil {
			log.Printf("fail: usecase.GetUsersByName, %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// usecaseから戻ってきたスライスのデータをJSON形式でフロントに返す
		// Go特有の型とJSON形式との変換はcontrollerの仕事！
		bytes, err := json.Marshal(users)
		if err != nil {
			log.Printf("fail: json.Marshal, %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(bytes)

	default:
		log.Printf("fail: HTTP Method is %s\n", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
