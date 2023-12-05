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
	// エンドポイントのルーター
	router := mux.NewRouter()

	// オプションハンドラー
	router.PathPrefix("/").HandlerFunc(controller.OptionsHandler).Methods("OPTIONS")

	//// テスト用ロボットの取得
	//router.HandleFunc("/robot", controller.SearchRobotHandler).Methods("GET")
	//// テスト用ロボットのポスト
	//router.HandleFunc("/robot", controller.RegisterRobotHandler).Methods("POST")
	//// テスト用ロボットの全取得
	//router.HandleFunc("/robots", controller.GetAllRobotsHandler).Methods("GET")

	// SignUpに合わせてユーザー情報をMySQLに保存
	router.HandleFunc("/signup", controller.UserSignUpHandler).Methods("POST")
	// ユーザー名の一意性の確認
	router.PathPrefix("/user/check/{user_name}").HandlerFunc(controller.UserCheckHandler).Methods("GET")
	// ユーザープロフィールの更新
	router.HandleFunc("/user/update", controller.UserUpdateHandler).Methods("PUT")
	// IDからユーザーデータを取得
	router.PathPrefix("/user/{user_name}").HandlerFunc(controller.UserGetHandler).Methods("GET")

	// ブログデータの作成
	router.HandleFunc("/blog", controller.BlogPostHandler).Methods("POST")
	// ブログIDをもとに詳細データを取得
	router.PathPrefix("/blog/{blog_id}").HandlerFunc(controller.BlogGetHandler).Methods("GET")
	// 逐次PUT
	router.HandleFunc("/blog/update", controller.BlogPutHandler).Methods("PUT")
	// 公開ボタンによるPUT
	router.HandleFunc("/blog/publish", controller.BlogPublishHandler).Methods("PUT")
	// ブログの削除
	router.PathPrefix("/blog/{blog_id}").HandlerFunc(controller.BlogDeleteHandler).Methods("DELETE")
	// 自分のブログを全取得
	router.PathPrefix("/blogs/draft/{user_id}").HandlerFunc(controller.BlogsGetHandler).Methods("GET")
	// ブログを検索して取得
	router.HandleFunc("/blogs/search", controller.BlogsSearchHandler).Methods("GET")
	// ブログを全て取得
	router.HandleFunc("/blogs/all", controller.BlogsGetAllHandler).Methods("GET")

	// 本データの作成
	router.HandleFunc("/book", controller.BookPostHandler).Methods("POST")
	// 本IDをもとに詳細データを取得
	router.PathPrefix("/book/{book_id}").HandlerFunc(controller.BookGetHandler).Methods("GET")
	// 逐次PUT
	router.HandleFunc("/book/update", controller.BookPutHandler).Methods("PUT")
	// 公開ボタンによるPUT
	router.HandleFunc("/book/publish", controller.BookPublishHandler).Methods("PUT")
	// 本の削除
	router.PathPrefix("/book/{book_id}").HandlerFunc(controller.BookDeleteHandler).Methods("DELETE")
	// 自分の本を全取得
	router.PathPrefix("/books/draft/{user_id}").HandlerFunc(controller.BooksGetHandler).Methods("GET")
	// 本を検索して取得
	router.HandleFunc("/books/search", controller.BooksSearchHandler).Methods("GET")
	// 本を全て取得
	router.HandleFunc("/books/all", controller.BooksGetAllHandler).Methods("GET")

	// 本のチャプターの作成
	router.HandleFunc("/chapter", controller.ChapterPostHandler).Methods("POST")
	// チャプターIDをもとに詳細データを取得
	router.PathPrefix("/chapter/{chapter_id}").HandlerFunc(controller.ChapterGetHandler).Methods("GET")
	// 参照中の本のチャプターを全取得
	router.PathPrefix("/chapters/{book_id}").HandlerFunc(controller.ChaptersGetHandler).Methods("GET")
	// 逐次PUT
	router.HandleFunc("/chapter/update", controller.ChapterPutHandler).Methods("PUT")

	// 動画データの作成
	router.HandleFunc("/video", controller.VideoPostHandler).Methods("POST")
	// 動画IDをもとに詳細データを取得
	router.PathPrefix("/video/{video_id}").HandlerFunc(controller.VideoGetHandler).Methods("GET")
	// 逐次PUT
	router.HandleFunc("/video/update", controller.VideoPutHandler).Methods("PUT")
	// 公開ボタンによるPUT
	router.HandleFunc("/video/publish", controller.VideoPublishHandler).Methods("PUT")
	// 動画の削除
	router.PathPrefix("/video/{video_id}").HandlerFunc(controller.VideoDeleteHandler).Methods("DELETE")
	// 自分の動画を全取得
	router.PathPrefix("/videos/draft/{user_id}").HandlerFunc(controller.VideosGetHandler).Methods("GET")
	// 動画を検索して取得
	router.HandleFunc("/videos/search", controller.VideosSearchHandler).Methods("GET")
	// 動画を全て取得
	router.HandleFunc("/videos/all", controller.VideosGetAllHandler).Methods("GET")

	// 作品データの作成
	router.HandleFunc("/work", controller.WorkPostHandler).Methods("POST")
	// 作品IDをもとに詳細データを取得
	router.PathPrefix("/work/{work_id}").HandlerFunc(controller.WorkGetHandler).Methods("GET")
	// 逐次PUT
	router.HandleFunc("/work/update", controller.WorkPutHandler).Methods("PUT")
	// 公開ボタンによるPUT
	router.HandleFunc("/work/publish", controller.WorkPublishHandler).Methods("PUT")
	// 作品の削除
	router.PathPrefix("/work/{work_id}").HandlerFunc(controller.WorkDeleteHandler).Methods("DELETE")
	// 自分の作品を全取得
	router.PathPrefix("/works/draft/{user_id}").HandlerFunc(controller.WorksGetHandler).Methods("GET")
	// 作品を検索して取得
	router.HandleFunc("/works/search", controller.WorksSearchHandler).Methods("GET")
	// 作品を全て取得
	router.HandleFunc("/works/all", controller.WorksGetAllHandler).Methods("GET")

	// コメントの作成
	router.HandleFunc("/comment", controller.CommentPostHandler).Methods("POST")
	// 参照中のアイテムのコメントを全取得
	router.PathPrefix("/comments/{item_category}/{item_id}").HandlerFunc(controller.CommentsGetHandler).Methods("GET")

	//// いいねのポスト
	//router.HandleFunc("/like", controller.LikePostHandler).Methods("POST")
	//// いいねの削除
	//router.PathPrefix("/like/{item_category}/{item_id}/{like_id}").HandlerFunc(controller.LikeDeleteHandler).Methods("DELETE")
	//// いいねの存在確認
	//router.PathPrefix("/like/{item_category}/{item_id}/{user_id}").HandlerFunc(controller.LikeCheckHandler).Methods("DELETE")

	http.Handle("/", router)

	// Ctrl+CでHTTPサーバー停止時にDBをクローズする
	closeDBWithSysCall()

	// 8080番ポートでリクエストを待ち受ける
	log.Println("Listening...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

// Ctrl+CでHTTPサーバー停止時にDBをクローズする
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
