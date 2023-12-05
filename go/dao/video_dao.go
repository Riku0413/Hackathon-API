package dao

import (
	"errors"
	"fmt"
	"github.com/Riku0413/Hackathon-API/model"
	"log"
	"net/url"
)

// トランザクションで新しいブログをポスト
func VideoPost(video model.Video) error {

	// ポストするデータが揃ったので、トランザクションを開始
	tx, err := Db.Begin()
	if err != nil {
		log.Printf("fail: Db.begin, %v\n", err)
		return err
	}

	// MySQLの操作
	_, err = tx.Exec("INSERT INTO video (id, title, url, user_id, birth_time, update_time, public, introduction) VALUES (?, ?, ?, ?, ?, ?, ?, ?)", video.Id, video.Title, "", video.UserId, video.BirthTime, video.UpdateTime, video.Public, "")
	if err != nil {
		// エラーが発生した場合、トランザクションをロールバックし、エラーを上位に返す
		tx.Rollback()
		log.Printf("fail: tx.Exec, %v\n", err)
		return err
	}

	// トランザクションを終了
	if err := tx.Commit(); err != nil {
		log.Printf("fail: tx.Commit, %v\n", err)
		return err
	}

	return nil
}

// トランザクションでブログを逐次更新
func VideoPut(video model.Video) error {

	// ポストするデータが揃ったので、トランザクションを開始
	tx, err := Db.Begin()
	if err != nil {
		log.Printf("fail: Db.begin, %v\n", err)
		return err
	}

	// MySQLの操作
	_, err = tx.Exec("UPDATE video SET title = ?, introduction = ?, url =?, update_time = ? WHERE id = ?", video.Title, video.Introduction, video.URL.String(), video.UpdateTime, video.Id)
	if err != nil {
		// エラーが発生した場合、トランザクションをロールバックし、エラーを上位に返す
		tx.Rollback()
		log.Printf("fail: tx.Exec, %v\n", err)
		return err
	}

	// トランザクションを終了
	if err := tx.Commit(); err != nil {
		log.Printf("fail: tx.Commit, %v\n", err)
		return err
	}

	return nil
}

// トランザクションでブログを公開 or 下書き保存
func VideoPublish(video model.Video) error {

	// ポストするデータが揃ったので、トランザクションを開始
	tx, err := Db.Begin()
	if err != nil {
		log.Printf("fail: Db.begin, %v\n", err)
		return err
	}

	// MySQLの操作
	_, err = tx.Exec("UPDATE video SET update_time = ?, public = ? WHERE id = ?", video.UpdateTime, video.Public, video.Id)
	if err != nil {
		// エラーが発生した場合、トランザクションをロールバックし、エラーを上位に返す
		tx.Rollback()
		log.Printf("fail: tx.Exec, %v\n", err)
		return err
	}

	// トランザクションを終了
	if err := tx.Commit(); err != nil {
		log.Printf("fail: tx.Commit, %v\n", err)
		return err
	}

	return nil
}

// idをもとにDBからデータを取得して、usecaseに返す
func VideoGet(id string) (model.Video, error) {
	// ここではトランザクションが不要！
	rows, err := Db.Query("SELECT * FROM video WHERE id = ?", id)
	if err != nil {
		log.Printf("fail: Db.Query, %v\n", err)
		return model.Video{}, err
	}
	defer rows.Close()

	// MySQL特有の型は、dao以外では登場させたくない → この処理はdaoで済ませる
	var v model.Video
	var urlString string
	if rows.Next() {
		if err := rows.Scan(&v.Id, &v.UserId, &v.Title, &urlString, &v.BirthTime, &v.UpdateTime, &v.Public, &v.Introduction); err != nil {
			log.Printf("fail: rows.Scan, %v\n", err)
			return model.Video{}, err
		}
		// URL文字列を*url.URLに変換
		v.URL, err = url.Parse(urlString)
		if err != nil {
			fmt.Println("URLのパースエラー:", err)
			return model.Video{}, err
		}
	} else {
		// クエリ結果が空の場合
		return model.Video{}, errors.New("Couldn't find the video specified")
	}

	return v, nil
}

// idをもとにDBからブログを削除して、削除したブログのidをusecaseに返す
func VideoDelete(id string) error {

	// 削除するブログのidが決まったので、トランザクションを開始
	tx, err := Db.Begin()
	if err != nil {
		log.Printf("fail: Db.begin, %v\n", err)
		return err
	}

	// データの存在性の確認
	var exists bool
	err = tx.QueryRow("SELECT EXISTS (SELECT 1 FROM video WHERE id = ?)", id).Scan(&exists)
	if err != nil {
		tx.Rollback()
		log.Printf("fail: tx.QueryRow, %v\n", err)
		return err
	}
	if !exists {
		tx.Rollback()
		return errors.New("Couldn't find the data of specified id")
	}

	// ブログを削除
	_, err = tx.Exec("DELETE FROM video WHERE id = ?", id)
	if err != nil {
		tx.Rollback()
		log.Printf("fail: tx.Exec, %v\n", err)
		return err
	}

	// トランザクションを終了
	if err := tx.Commit(); err != nil {
		log.Printf("fail: tx.Commit, %v\n", err)
		return err
	}

	return nil
}

// user_idに紐づくブログを最終更新日時の降順で取得
func VideosGet(userId string) ([]model.Video, error) {
	rows, err := Db.Query("SELECT id, title, update_time FROM video WHERE user_id = ? ORDER BY update_time DESC", userId)
	if err != nil {
		log.Printf("fail: Db.Query, %v\n", err)
		return nil, err
	}

	// rowsはここで、しっかり処理して、Go特有の型に変換してからusecaseに戻す
	videos := make([]model.Video, 0)
	for rows.Next() {
		var v model.Video
		// ブログID、タイトル、最終更新日時、の3つだけを返す！
		if err := rows.Scan(&v.Id, &v.Title, &v.UpdateTime); err != nil {
			log.Printf("fail: rows.Scan, %v\n", err)
			if err := rows.Close(); err != nil { // 500を返して終了するが、その前にrowsのClose処理が必要
				log.Printf("fail: rows.Close(), %v\n", err)
			}
			return nil, err
		}
		videos = append(videos, v) // users にデータを順々に格納！
	}

	return videos, nil
}

// 検索キーから動画を最終更新日時の降順で取得
func VideosSearch(key string) ([]model.Video, error) {
	query := "SELECT id, title, birth_time, update_time FROM video WHERE (title LIKE ?) AND public = 1 ORDER BY update_time DESC"
	keyword := "%" + key + "%" // キーワードを含む部分文字列を作成
	rows, err := Db.Query(query, keyword)
	if err != nil {
		log.Printf("fail: Db.Query, %v\n", err)
		return nil, err
	}

	// rowsはここで、しっかり処理して、Go特有の型に変換してからusecaseに戻す
	videos := make([]model.Video, 0)
	for rows.Next() {
		var v model.Video
		// 動画ID、タイトル、最終更新日時、の3つだけを返す！
		if err := rows.Scan(&v.Id, &v.Title, &v.BirthTime, &v.UpdateTime); err != nil {
			log.Printf("fail: rows.Scan, %v\n", err)
			if err := rows.Close(); err != nil { // 500を返して終了するが、その前にrowsのClose処理が必要
				log.Printf("fail: rows.Close(), %v\n", err)
			}
			return nil, err
		}
		videos = append(videos, v)
	}

	return videos, nil
}

// 全ての動画を取得
func VideosGetAll() ([]model.Video, error) {
	rows, err := Db.Query("SELECT id, title, birth_time, update_time FROM video WHERE public = 1")
	if err != nil {
		log.Printf("fail: Db.Query, %v\n", err)
		return nil, err
	}

	videos := make([]model.Video, 0)
	for rows.Next() {
		var v model.Video
		if err := rows.Scan(&v.Id, &v.Title, &v.BirthTime, &v.UpdateTime); err != nil {
			log.Printf("fail: rows.Scan, %v\n", err)
			if err := rows.Close(); err != nil {
				log.Printf("fail: rows.Close(), %v\n", err)
			}
			return nil, err
		}
		videos = append(videos, v)
	}

	return videos, nil
}
