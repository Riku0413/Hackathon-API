package main

import (
	"github.com/Riku0413/Hackathon-API/controller"
	"github.com/Riku0413/Hackathon-API/dao"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
)

func main() {
	// /userのエンドポイント　GETとPOSTで分岐させる
	router := mux.NewRouter();
	router.HandleFunc("/user", controller.OptionsHandler).Methods("OPTIONS");
	router.HandleFunc("/user", controller.SearchUserHandler).Methods("GET");
	router.HandleFunc("/user", controller.RegisterUserHandler).Methods("POST");
	// http.Handle("/", router);

	// /usersのエンドポイント
	// http.HandleFunc("/users", controller.GetAllUsersHandler)

	router.HandleFunc("/users", controller.OptionsHandler).Methods("OPTIONS");
	router.HandleFunc("/users", controller.GetAllUsersHandler).Methods("GET");

	http.Handle("/", router);

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

		if err := dao.Db.Close(); err != nil {
			log.Fatal(err)
		}
		log.Printf("success: dao.Db.Close()")
		os.Exit(0)
	}()
}
