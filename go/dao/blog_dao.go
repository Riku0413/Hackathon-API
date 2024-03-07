package dao

import (
	"errors"
	"github.com/Riku0413/Hackathon-API/model"
	"log"
)

// トランザクションで新しいブログをポスト
func BlogPost(blog model.Blog) error {

	// ポストするデータが揃ったので、トランザクションを開始
	tx, err := Db.Begin()
	if err != nil {
		log.Printf("fail: Db.begin, %v\n", err)
		return err
	}

	// MySQLの操作
	_, err = tx.Exec("INSERT INTO blog (id, title, content, user_id, birth_time, update_time, public, likes) VALUES (?, ?, ?, ?, ?, ?, ?, ?)", blog.Id, blog.Title, blog.Content, blog.UserId, blog.BirthTime, blog.UpdateTime, blog.Public, 0)
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
func BlogPut(blog model.Blog) error {

	// ポストするデータが揃ったので、トランザクションを開始
	tx, err := Db.Begin()
	if err != nil {
		log.Printf("fail: Db.begin, %v\n", err)
		return err
	}

	// MySQLの操作
	_, err = tx.Exec("UPDATE blog SET title = ?, content = ?, update_time = ? WHERE id = ?", blog.Title, blog.Content, blog.UpdateTime, blog.Id)
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
func BlogPublish(blog model.Blog) error {

	// ポストするデータが揃ったので、トランザクションを開始
	tx, err := Db.Begin()
	if err != nil {
		log.Printf("fail: Db.begin, %v\n", err)
		return err
	}

	// MySQLの操作
	_, err = tx.Exec("UPDATE blog SET update_time = ?, public = ? WHERE id = ?", blog.UpdateTime, blog.Public, blog.Id)
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
func BlogGet(id string) (model.Blog, error) {
	// ここではトランザクションが不要！
	rows, err := Db.Query("SELECT * FROM blog WHERE id = ?", id)
	if err != nil {
		log.Printf("fail: Db.Query, %v\n", err)
		return model.Blog{}, err
	}
	defer rows.Close()

	// MySQL特有の型は、dao以外では登場させたくない → この処理はdaoで済ませる
	var b model.Blog
	if rows.Next() {
		if err := rows.Scan(&b.Id, &b.UserId, &b.Title, &b.Content, &b.BirthTime, &b.UpdateTime, &b.Public, &b.Likes); err != nil {
			log.Printf("fail: rows.Scan, %v\n", err)
			return model.Blog{}, err
		}
	} else {
		// クエリ結果が空の場合
		return model.Blog{}, errors.New("Couldn't find the blog specified")
	}

	return b, nil
}

// idをもとにDBからブログを削除して、削除したブログのidをusecaseに返す
func BlogDelete(id string) error {

	// 削除するブログのidが決まったので、トランザクションを開始
	tx, err := Db.Begin()
	if err != nil {
		log.Printf("fail: Db.begin, %v\n", err)
		return err
	}

	// データの存在性の確認
	var exists bool
	err = tx.QueryRow("SELECT EXISTS (SELECT 1 FROM blog WHERE id = ?)", id).Scan(&exists)
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
	_, err = tx.Exec("DELETE FROM blog WHERE id = ?", id)
	if err != nil {
		tx.Rollback()
		log.Printf("fail: tx.Exec, %v\n", err)
		return err
	}

	// 関連するいいねを削除
	_, err = tx.Exec("DELETE FROM likes_blog WHERE blog_id = ?", id)
	if err != nil {
		tx.Rollback()
		log.Printf("fail: tx.Exec, %v\n", err)
		return err
	}

	// 関連するコメントを削除
	_, err = tx.Exec("DELETE FROM comment_blog WHERE blog_id = ?", id)
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
func BlogsGet(userId string) ([]model.Blog, error) {
	rows, err := Db.Query("SELECT id, title, update_time, likes FROM blog WHERE user_id = ? ORDER BY update_time DESC", userId)
	if err != nil {
		log.Printf("fail: Db.Query, %v\n", err)
		return nil, err
	}

	// rowsはここで、しっかり処理して、Go特有の型に変換してからusecaseに戻す
	blogs := make([]model.Blog, 0)
	for rows.Next() {
		var b model.Blog
		// ブログID、タイトル、最終更新日時、の3つだけを返す！
		if err := rows.Scan(&b.Id, &b.Title, &b.UpdateTime, &b.Likes); err != nil {
			log.Printf("fail: rows.Scan, %v\n", err)
			if err := rows.Close(); err != nil { // 500を返して終了するが、その前にrowsのClose処理が必要
				log.Printf("fail: rows.Close(), %v\n", err)
			}
			return nil, err
		}
		blogs = append(blogs, b) // users にデータを順々に格納！
	}

	return blogs, nil
}

// 検索キーからブログを最終更新日時の降順で取得
func BlogsSearch(key string) ([]model.Blog, error) {
	query := "SELECT id, title, birth_time, update_time, likes FROM blog WHERE (title LIKE ? OR content LIKE ?) AND public = 1 ORDER BY update_time DESC"
	keyword := "%" + key + "%" // キーワードを含む部分文字列を作成
	rows, err := Db.Query(query, keyword, keyword)
	if err != nil {
		log.Printf("fail: Db.Query, %v\n", err)
		return nil, err
	}

	// rowsはここで、しっかり処理して、Go特有の型に変換してからusecaseに戻す
	blogs := make([]model.Blog, 0)
	for rows.Next() {
		var b model.Blog
		// ブログID、タイトル、最終更新日時、の3つだけを返す！
		if err := rows.Scan(&b.Id, &b.Title, &b.BirthTime, &b.UpdateTime, &b.Likes); err != nil {
			log.Printf("fail: rows.Scan, %v\n", err)
			if err := rows.Close(); err != nil { // 500を返して終了するが、その前にrowsのClose処理が必要
				log.Printf("fail: rows.Close(), %v\n", err)
			}
			return nil, err
		}
		blogs = append(blogs, b) // users にデータを順々に格納！
	}

	return blogs, nil
}

// 全てのブログを取得
func BlogsGetAll() ([]model.Blog, error) {
	rows, err := Db.Query("SELECT id, title, birth_time, update_time, likes FROM blog WHERE public = 1")
	if err != nil {
		log.Printf("fail: Db.Query, %v\n", err)
		return nil, err
	}

	blogs := make([]model.Blog, 0)
	for rows.Next() {
		var b model.Blog
		if err := rows.Scan(&b.Id, &b.Title, &b.BirthTime, &b.UpdateTime, &b.Likes); err != nil {
			log.Printf("fail: rows.Scan, %v\n", err)
			if err := rows.Close(); err != nil {
				log.Printf("fail: rows.Close(), %v\n", err)
			}
			return nil, err
		}
		blogs = append(blogs, b)
	}

	return blogs, nil
}
