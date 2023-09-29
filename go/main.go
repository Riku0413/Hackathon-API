// フロントからのリクエストを受け取り、MySQLに繋げるためのサーバー
// 全ユーザーの取得
// ユーザーのDBへの登録 + 全ユーザーの取得

package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/oklog/ulid/v2"
	_ "github.com/oklog/ulid/v2" // ULIDをimport
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/rs/cors"

	// "net/url"
)

type UserResForHTTPGet struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// ① GoプログラムからMySQLへ接続
var db *sql.DB

func init() {
	// ①-1

	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlPwd := os.Getenv("MYSQL_PASS")
	mysqlHost := os.Getenv("MYSQL_HOST")
	mysqlDatabase := os.Getenv("MYSQL_DATABASE")

	// mysqlUser := "uttc"
	// mysqlPwd := "rad32first"
	// mysqlHost := "34.42.70.51"
	// mysqlPort := "3306" // ポート番号はMySQLのデフォルトポート（3306）に設定します
	// mysqlDatabase := "unix(/cloudsql/term4-riku-kobayashi:us-central1:uttc)"
	// mysqlDatabase = url.QueryEscape(mysqlDatabase)

	// connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", mysqlUser, mysqlPwd, mysqlHost, mysqlPort, mysqlDatabase)
	// _db, err := sql.Open("mysql", connStr)

	// ①-2
	connStr := fmt.Sprintf("%s:%s@%s/%s", mysqlUser, mysqlPwd, mysqlHost, mysqlDatabase)
	_db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatalf("fail: sql.Open, %v\n", err)
	}
	// ①-3
	if err := _db.Ping(); err != nil {
		log.Fatalf("fail: _db.Ping, %v\n", err)
	}
	db = _db
}

func handler(w http.ResponseWriter, r *http.Request) {
	// アクセス権を追加
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

	switch r.Method {
	// ② /userでリクエストされたらnameパラメーターと一致する名前を持つレコードをJSON形式で返す
	case http.MethodGet:
		// ②-1
		queryParams := r.URL.Query()
		name := queryParams.Get("name")
		if name == "" {
			log.Println("fail: name is empty")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// ②-2
		rows, err := db.Query("SELECT id, name, age FROM user WHERE name = ?", name)
		if err != nil {
			log.Printf("fail: db.Query, %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// ②-3
		users := make([]UserResForHTTPGet, 0)
		for rows.Next() {
			var u UserResForHTTPGet
			if err := rows.Scan(&u.Id, &u.Name, &u.Age); err != nil {
				log.Printf("fail: rows.Scan, %v\n", err)

				if err := rows.Close(); err != nil { // 500を返して終了するが、その前にrowsのClose処理が必要
					log.Printf("fail: rows.Close(), %v\n", err)
				}
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			users = append(users, u)
		}

		// ②-4
		bytes, err := json.Marshal(users)
		if err != nil {
			log.Printf("fail: json.Marshal, %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(bytes)

	// POSTリクエストにより、名前と年齢の情報をデータベースに追加する
	case http.MethodPost:
		decoder := json.NewDecoder(r.Body)
		var newUser UserResForHTTPGet
		if err := decoder.Decode(&newUser); err != nil {
			log.Printf("fail: json decode, %v\n", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		// name, age の条件を満たしているかチェック
		if newUser.Name == "" || len(newUser.Name) > 50 {
			log.Println("fail: name is invalid")
			w.WriteHeader(http.StatusBadRequest)
			return
		} else if newUser.Age < 20 || newUser.Age > 80 {
			log.Println("fail: age is invalid")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// newUser の内容をクライアントに出力（デバッグ用）
		//responseMsg := fmt.Sprintf("Received new user: %+v", newUser)
		//w.Write([]byte(responseMsg))

		// ULIDを用いて、乱数で採番
		entropy := rand.New(rand.NewSource(time.Now().UnixNano()))
		ms := ulid.Timestamp(time.Now())
		ulidValue, err := ulid.New(ms, entropy)
		ulidString := ulidValue.String() // ULIDを文字列に変換

		if err != nil {
			log.Printf("fail: ulid generate, %v\n", err)
			w.Write([]byte("error number : 1")) // エラーメッセージをクライアントに返す（デバッグ用）
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// MySQLの操作
		_, err = db.Exec("INSERT INTO user (id, name, age) VALUES (?, ?, ?)", ulidString, newUser.Name, newUser.Age)
		if err != nil {
			log.Printf("fail: db.Exec, %v\n", err)
			w.Write([]byte("error number : 2")) // エラーメッセージをクライアントに返す（デバッグ用）
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// クライアントへのレスポンスを作成
		response := map[string]interface{}{
			"id": ulidString,
		}
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			log.Printf("fail: json.Marshal, %v\n", err)
			w.Write([]byte("error number : 3")) // エラーメッセージをクライアントに返す（デバッグ用）
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
		w.WriteHeader(http.StatusOK)
		log.Println("HTTP Status Code 200")

	default:
		log.Printf("fail: HTTP Method is %s\n", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

// 全データの取得処理
func handler_2(w http.ResponseWriter, r *http.Request) {
	// アクセス権を追加
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Headers", "*") // ここのスペルミス！！！
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

	switch r.Method {
	case http.MethodOptions:
		log.Printf("options")
		w.WriteHeader(http.StatusOK)
		return

	case http.MethodGet:
		log.Printf("get")
		rows, err := db.Query("SELECT id, name, age FROM user")
		if err != nil {
			log.Printf("fail: db.Query, %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		users := make([]UserResForHTTPGet, 0)
		for rows.Next() {
			var u UserResForHTTPGet
			if err := rows.Scan(&u.Id, &u.Name, &u.Age); err != nil {
				log.Printf("fail: rows.Scan, %v\n", err)

				if err := rows.Close(); err != nil { // 500を返して終了するが、その前にrowsのClose処理が必要
					log.Printf("fail: rows.Close(), %v\n", err)
				}
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			users = append(users, u) // users にデータを順々に格納！
		}

		bytes, err := json.Marshal(users)
		if err != nil {
			log.Printf("fail: json.Marshal, %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(bytes)

	// POSTリクエストにより、名前と年齢の情報をデータベースに追加する
	case http.MethodPost:
		decoder := json.NewDecoder(r.Body)
		var newUser UserResForHTTPGet
		if err := decoder.Decode(&newUser); err != nil {
			log.Printf("fail: json decode, %v\n", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		// name, age の条件を満たしているかチェック
		if newUser.Name == "" || len(newUser.Name) > 50 {
			log.Println("fail: name is invalid")
			w.WriteHeader(http.StatusBadRequest)
			return
		} else if newUser.Age < 20 || newUser.Age > 80 {
			log.Println("fail: age is invalid")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// ULIDを用いて、乱数で採番
		entropy := rand.New(rand.NewSource(time.Now().UnixNano()))
		ms := ulid.Timestamp(time.Now())
		ulidValue, err := ulid.New(ms, entropy)
		ulidString := ulidValue.String() // ULIDを文字列に変換

		if err != nil {
			log.Printf("fail: ulid generate, %v\n", err)
			w.Write([]byte("error number : 1")) // エラーメッセージをクライアントに返す（デバッグ用）
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// MySQLの操作
		_, err = db.Exec("INSERT INTO user (id, name, age) VALUES (?, ?, ?)", ulidString, newUser.Name, newUser.Age)
		if err != nil {
			log.Printf("fail: db.Exec, %v\n", err)
			w.Write([]byte("error number : 2")) // エラーメッセージをクライアントに返す（デバッグ用）
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// クライアントへのレスポンスを作成
		response := map[string]interface{}{
			"id": ulidString,
		}
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			log.Printf("fail: json.Marshal, %v\n", err)
			w.Write([]byte("error number : 3")) // エラーメッセージをクライアントに返す（デバッグ用）
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
		w.WriteHeader(http.StatusOK)
		log.Println("HTTP Status Code 200")

	default:
		log.Printf("fail: HTTP Method is %s\n", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func main() {

	// ② /userでリクエストされたらnameパラメーターと一致する名前を持つレコードをJSON形式で返す
	http.HandleFunc("/user", handler)

	// 追加のエンドポイント
	http.HandleFunc("/users", handler_2)

	// ③ Ctrl+CでHTTPサーバー停止時にDBをクローズする
	closeDBWithSysCall()

	// 8000番ポートでリクエストを待ち受ける
	log.Println("Listening...")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

// ③ Ctrl+CでHTTPサーバー停止時にDBをクローズする
func closeDBWithSysCall() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		s := <-sig
		log.Printf("received syscall, %v", s)

		if err := db.Close(); err != nil {
			log.Fatal(err)
		}
		log.Printf("success: db.Close()")
		os.Exit(0)
	}()
}
