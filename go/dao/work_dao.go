package dao

import (
	"errors"
	"github.com/Riku0413/Hackathon-API/model"
	"log"
)

// トランザクションで新しいブログをポスト
func WorkPost(work model.Work) error {

	// ポストするデータが揃ったので、トランザクションを開始
	tx, err := Db.Begin()
	if err != nil {
		log.Printf("fail: Db.begin, %v\n", err)
		return err
	}

	// MySQLの操作
	_, err = tx.Exec("INSERT INTO work (id, title, url, user_id, birth_time, update_time, public, introduction) VALUES (?, ?, ?, ?, ?, ?, ?, ?)", work.Id, work.Title, "", work.UserId, work.BirthTime, work.UpdateTime, work.Public, "")
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
func WorkPut(work model.Work) error {

	// ポストするデータが揃ったので、トランザクションを開始
	tx, err := Db.Begin()
	if err != nil {
		log.Printf("fail: Db.begin, %v\n", err)
		return err
	}

	// MySQLの操作
	_, err = tx.Exec("UPDATE work SET title = ?, introduction = ?, url =?, update_time = ? WHERE id = ?", work.Title, work.Introduction, work.URL, work.UpdateTime, work.Id)
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
func WorkPublish(work model.Work) error {

	// ポストするデータが揃ったので、トランザクションを開始
	tx, err := Db.Begin()
	if err != nil {
		log.Printf("fail: Db.begin, %v\n", err)
		return err
	}

	// MySQLの操作
	_, err = tx.Exec("UPDATE work SET update_time = ?, public = ? WHERE id = ?", work.UpdateTime, work.Public, work.Id)
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
func WorkGet(id string) (model.Work, error) {
	// ここではトランザクションが不要！
	rows, err := Db.Query("SELECT * FROM work WHERE id = ?", id)
	if err != nil {
		log.Printf("fail: Db.Query, %v\n", err)
		return model.Work{}, err
	}
	defer rows.Close()

	// MySQL特有の型は、dao以外では登場させたくない → この処理はdaoで済ませる
	var w model.Work
	//var urlString string
	if rows.Next() {
		if err := rows.Scan(&w.Id, &w.UserId, &w.Title, &w.Introduction, &w.URL, &w.BirthTime, &w.UpdateTime, &w.Public); err != nil {
			log.Printf("fail: rows.Scan, %v\n", err)
			return model.Work{}, err
		}
		// URL文字列を*url.URLに変換
		//w.URL, err = url.Parse(urlString)
		//if err != nil {
		//	fmt.Println("URLのパースエラー:", err)
		//	return model.Work{}, err
		//}
	} else {
		// クエリ結果が空の場合
		return model.Work{}, errors.New("Couldn't find the work specified")
	}

	return w, nil
}

// idをもとにDBからブログを削除して、削除したブログのidをusecaseに返す
func WorkDelete(id string) error {

	// 削除するブログのidが決まったので、トランザクションを開始
	tx, err := Db.Begin()
	if err != nil {
		log.Printf("fail: Db.begin, %v\n", err)
		return err
	}

	// データの存在性の確認
	var exists bool
	err = tx.QueryRow("SELECT EXISTS (SELECT 1 FROM work WHERE id = ?)", id).Scan(&exists)
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
	_, err = tx.Exec("DELETE FROM work WHERE id = ?", id)
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
func WorksGet(userId string) ([]model.Work, error) {
	rows, err := Db.Query("SELECT id, title, update_time FROM work WHERE user_id = ? ORDER BY update_time DESC", userId)
	if err != nil {
		log.Printf("fail: Db.Query, %v\n", err)
		return nil, err
	}

	// rowsはここで、しっかり処理して、Go特有の型に変換してからusecaseに戻す
	works := make([]model.Work, 0)
	for rows.Next() {
		var w model.Work
		// ブログID、タイトル、最終更新日時、の3つだけを返す！
		if err := rows.Scan(&w.Id, &w.Title, &w.UpdateTime); err != nil {
			log.Printf("fail: rows.Scan, %v\n", err)
			if err := rows.Close(); err != nil { // 500を返して終了するが、その前にrowsのClose処理が必要
				log.Printf("fail: rows.Close(), %v\n", err)
			}
			return nil, err
		}
		works = append(works, w) // users にデータを順々に格納！
	}

	return works, nil
}

// 検索キーから作品を最終更新日時の降順で取得
func WorksSearch(key string) ([]model.Work, error) {
	query := "SELECT id, title, birth_time, update_time FROM work WHERE (title LIKE ?) AND public = 1 ORDER BY update_time DESC"
	keyword := "%" + key + "%" // キーワードを含む部分文字列を作成
	rows, err := Db.Query(query, keyword)
	if err != nil {
		log.Printf("fail: Db.Query, %v\n", err)
		return nil, err
	}

	// rowsはここで、しっかり処理して、Go特有の型に変換してからusecaseに戻す
	works := make([]model.Work, 0)
	for rows.Next() {
		var w model.Work
		// 動画ID、タイトル、最終更新日時、の3つだけを返す！
		if err := rows.Scan(&w.Id, &w.Title, &w.BirthTime, &w.UpdateTime); err != nil {
			log.Printf("fail: rows.Scan, %v\n", err)
			if err := rows.Close(); err != nil { // 500を返して終了するが、その前にrowsのClose処理が必要
				log.Printf("fail: rows.Close(), %v\n", err)
			}
			return nil, err
		}
		works = append(works, w)
	}

	return works, nil
}

// 全ての動画を取得
func WorksGetAll() ([]model.Work, error) {
	rows, err := Db.Query("SELECT id, title, birth_time, update_time FROM work WHERE public = 1")
	if err != nil {
		log.Printf("fail: Db.Query, %v\n", err)
		return nil, err
	}

	works := make([]model.Work, 0)
	for rows.Next() {
		var w model.Work
		if err := rows.Scan(&w.Id, &w.Title, &w.BirthTime, &w.UpdateTime); err != nil {
			log.Printf("fail: rows.Scan, %v\n", err)
			if err := rows.Close(); err != nil {
				log.Printf("fail: rows.Close(), %v\n", err)
			}
			return nil, err
		}
		works = append(works, w)
	}

	return works, nil
}
