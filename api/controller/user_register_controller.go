package controller

import (
	"encoding/json"
	"github.com/Riku0413/Hackathon-API/model"
	"github.com/Riku0413/Hackathon-API/usecase"
	"github.com/oklog/ulid/v2"
	"log"
	"math/rand"
	"net/http"
	"time"
)

// 新しいユーザーデータのポスト
func RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	// ここいるのかな
	w.Header().Set("Access-Control-Allow-Origin", "*") // ここの条件は審議！
	w.Header().Set("Access-Control-Allow-Headers", "*") // ここのスペルミス！！！
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	
	switch r.Method {
	// case http.MethodOptions:
	// 	log.Printf("options 2") // これいる？　いるなら統一しよう
	// 	w.WriteHeader(http.StatusOK)
	// 	return
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
		if err := model.User.CheckName(newUser); err != nil {
			log.Printf("fail: model.User.CheckName, %v\n", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		} else if err := model.User.CheckAge(newUser); err != nil {
			log.Printf("fail: model.User.CheckAge, %v\n", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// ULIDを用いて、乱数で採番
		entropy := rand.New(rand.NewSource(time.Now().UnixNano()))
		ms := ulid.Timestamp(time.Now())
		ulidValue, err := ulid.New(ms, entropy) // ここら辺でエラーになることはないの？？
		if err != nil {
			log.Printf("fail: ulid generate, %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		ulidString := ulidValue.String() // ULIDを文字列に変換　// ここでエラーになることは？
		newUser.Id = ulidString

		// ここまでで、ポストするデータの妥当性チェックと、そのデータの、Goの型での準備が完了！

		// usecaseにバトンタッチし、error or not だけを結果として受け取る！
		if err := usecase.RegisterUser(newUser); err != nil {
			log.Printf("fail: usecase.RegisterUser, %v\n", err) // エラー表示はこの書き方で統一する
			w.WriteHeader(http.StatusInternalServerError)       // ここら辺これでいいのかな
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
