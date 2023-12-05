package controller

import (
	"encoding/json"
	"github.com/Riku0413/Hackathon-API/model"
	"github.com/Riku0413/Hackathon-API/usecase"
	"log"
	"net/http"
)

// 新しいユーザーデータのポスト
func UserSignUpHandler(w http.ResponseWriter, r *http.Request) {
	// ここいるのかな
	w.Header().Set("Access-Control-Allow-Origin", "*") // ここの条件は審議！
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")

	switch r.Method {
	case http.MethodPost:
		// ポストするJSON形式データをデコードしてGoの形式に変換
		decoder := json.NewDecoder(r.Body)
		var newUser model.User
		if err := decoder.Decode(&newUser); err != nil {
			log.Printf("fail: json decode, %v\n", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		defer r.Body.Close() // ここはエラーにならないの？？？

		// name, age の条件を満たしているかチェック（バリデーション）
		if err := model.User.UserCheckUID(newUser); err != nil {
			log.Printf("fail: model.User.CheckHuman, %v\n", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// ここまでで、ポストするデータの妥当性チェックと、そのデータの、Goの型での準備が完了！

		// usecaseにバトンタッチし、error or not だけを結果として受け取る！
		if err := usecase.UserSignUp(newUser); err != nil {
			log.Printf("fail: usecase.HumanSignup, %v\n", err) // エラー表示はこの書き方で統一する
			w.WriteHeader(http.StatusInternalServerError)      // ここら辺これでいいのかな
			return
		}

		// クライアントへのレスポンスを作成し、JSON形式に変換
		// JSON形式はcontrollerでしか登場させない！
		response := map[string]interface{}{
			"id": newUser.Id, // ここのエラーはないの？
		}
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			log.Printf("fail: json.Marshal, %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// フロントにレスポンスとして登録したIDを返す
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
		// w.WriteHeader(http.StatusOK) // ここは要らない！負の遺産！
		log.Printf("post")

	default:
		log.Printf("fail: HTTP Method is %s\n", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
