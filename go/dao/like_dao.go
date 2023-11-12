package dao

import (
	"database/sql"
	"errors"
	"github.com/Riku0413/Hackathon-API/model"
	"log"
)

func LikePost(like model.Like) error {

	// ポストするデータが揃ったので、トランザクションを開始
	tx, err := Db.Begin()
	if err != nil {
		log.Printf("fail: Db.begin, %v\n", err)
		return err
	}

	// MySQLの操作
	if like.ItemCategory == "blog" {
		_, err = tx.Exec("INSERT INTO likes_blog (id, user_id, blog_id, birth_time) VALUES (?, ?, ?, ?)", like.Id, like.UserId, like.ItemId, like.BirthTime)
		if err != nil {
			tx.Rollback()
			log.Printf("fail: tx.Exec, %v\n", err)
			return err
		}
		_, err = tx.Exec("UPDATE blog SET likes = likes + 1 WHERE id = ?", like.ItemId)
		if err != nil {
			tx.Rollback()
			log.Printf("fail: tx.Exec, %v\n", err)
			return err
		}
	} else if like.ItemCategory == "book" {
		_, err = tx.Exec("INSERT INTO likes_book (id, user_id, book_id, birth_time) VALUES (?, ?, ?, ?)", like.Id, like.UserId, like.ItemId, like.BirthTime)
		if err != nil {
			tx.Rollback()
			log.Printf("fail: tx.Exec, %v\n", err)
			return err
		}
		_, err = tx.Exec("UPDATE book SET likes = likes + 1 WHERE id = ?", like.ItemId)
		if err != nil {
			tx.Rollback()
			log.Printf("fail: tx.Exec, %v\n", err)
			return err
		}
	} else if like.ItemCategory == "video" {
		_, err = tx.Exec("INSERT INTO likes_video (id, user_id, video_id, birth_time) VALUES (?, ?, ?, ?)", like.Id, like.UserId, like.ItemId, like.BirthTime)
		if err != nil {
			tx.Rollback()
			log.Printf("fail: tx.Exec, %v\n", err)
			return err
		}
		_, err = tx.Exec("UPDATE video SET likes = likes + 1 WHERE id = ?", like.ItemId)
		if err != nil {
			tx.Rollback()
			log.Printf("fail: tx.Exec, %v\n", err)
			return err
		}
	} else if like.ItemCategory == "work" {
		_, err = tx.Exec("INSERT INTO likes_work (id, user_id, work_id, birth_time) VALUES (?, ?, ?, ?)", like.Id, like.UserId, like.ItemId, like.BirthTime)
		if err != nil {
			tx.Rollback()
			log.Printf("fail: tx.Exec, %v\n", err)
			return err
		}
		_, err = tx.Exec("UPDATE work SET likes = likes + 1 WHERE id = ?", like.ItemId)
		if err != nil {
			tx.Rollback()
			log.Printf("fail: tx.Exec, %v\n", err)
			return err
		}
	} else {
		return errors.New("Item category is invalid")
	}

	// トランザクションを終了
	if err := tx.Commit(); err != nil {
		log.Printf("fail: tx.Commit, %v\n", err)
		return err
	}

	return nil
}

// トランザクションでブログを逐次更新
//func BlogPut(blog model.Blog) error {
//
//	// ポストするデータが揃ったので、トランザクションを開始
//	tx, err := Db.Begin()
//	if err != nil {
//		log.Printf("fail: Db.begin, %v\n", err)
//		return err
//	}
//
//	// MySQLの操作
//	_, err = tx.Exec("UPDATE blog SET title = ?, content = ?, update_time = ? WHERE id = ?", blog.Title, blog.Content, blog.UpdateTime, blog.Id)
//	if err != nil {
//		// エラーが発生した場合、トランザクションをロールバックし、エラーを上位に返す
//		tx.Rollback()
//		log.Printf("fail: tx.Exec, %v\n", err)
//		return err
//	}
//
//	// トランザクションを終了
//	if err := tx.Commit(); err != nil {
//		log.Printf("fail: tx.Commit, %v\n", err)
//		return err
//	}
//
//	return nil
//}

// トランザクションでブログを公開 or 下書き保存
//func BlogPublish(blog model.Blog) error {
//
//	// ポストするデータが揃ったので、トランザクションを開始
//	tx, err := Db.Begin()
//	if err != nil {
//		log.Printf("fail: Db.begin, %v\n", err)
//		return err
//	}
//
//	// MySQLの操作
//	_, err = tx.Exec("UPDATE blog SET update_time = ?, public = ? WHERE id = ?", blog.UpdateTime, blog.Public, blog.Id)
//	if err != nil {
//		// エラーが発生した場合、トランザクションをロールバックし、エラーを上位に返す
//		tx.Rollback()
//		log.Printf("fail: tx.Exec, %v\n", err)
//		return err
//	}
//
//	// トランザクションを終了
//	if err := tx.Commit(); err != nil {
//		log.Printf("fail: tx.Commit, %v\n", err)
//		return err
//	}
//
//	return nil
//}

// idをもとにDBからデータを取得して、usecaseに返す
//func BlogGet(id string) (model.Blog, error) {
//	// ここではトランザクションが不要！
//	rows, err := Db.Query("SELECT * FROM blog WHERE id = ?", id)
//	if err != nil {
//		log.Printf("fail: Db.Query, %v\n", err)
//		return model.Blog{}, err
//	}
//	defer rows.Close()
//
//	// MySQL特有の型は、dao以外では登場させたくない → この処理はdaoで済ませる
//	var b model.Blog
//	if rows.Next() {
//		if err := rows.Scan(&b.Id, &b.Title, &b.Content, &b.UserId, &b.BirthTime, &b.UpdateTime, &b.Public); err != nil {
//			log.Printf("fail: rows.Scan, %v\n", err)
//			return model.Blog{}, err
//		}
//	} else {
//		// クエリ結果が空の場合
//		return model.Blog{}, errors.New("Couldn't find the blog specified")
//	}
//
//	return b, nil
//}

// idをもとにDBからブログを削除して、削除したブログのidをusecaseに返す
func LikeDelete(id string, category string, itemId string) error {

	// 削除するブログのidが決まったので、トランザクションを開始
	tx, err := Db.Begin()
	if err != nil {
		log.Printf("fail: Db.begin, %v\n", err)
		return err
	}

	if category == "blog" {
		var exists bool
		err = tx.QueryRow("SELECT EXISTS (SELECT 1 FROM likes_blog WHERE id = ?)", id).Scan(&exists)
		if err != nil {
			tx.Rollback()
			log.Printf("fail: tx.QueryRow, %v\n", err)
			return err
		}
		if !exists {
			tx.Rollback()
			return errors.New("Couldn't find the data of specified id")
		}

		_, err = tx.Exec("DELETE FROM likes_blog WHERE id = ?", id)
		if err != nil {
			tx.Rollback()
			log.Printf("fail: tx.Exec, %v\n", err)
			return err
		}

		_, err = tx.Exec("UPDATE blog SET likes = likes - 1 WHERE id = ?", itemId)
		if err != nil {
			tx.Rollback()
			log.Printf("fail: tx.Exec, %v\n", err)
			return err
		}
	} else if category == "book" {
		var exists bool
		err = tx.QueryRow("SELECT EXISTS (SELECT 1 FROM likes_book WHERE id = ?)", id).Scan(&exists)
		if err != nil {
			tx.Rollback()
			log.Printf("fail: tx.QueryRow, %v\n", err)
			return err
		}
		if !exists {
			tx.Rollback()
			return errors.New("Couldn't find the data of specified id")
		}

		_, err = tx.Exec("DELETE FROM likes_book WHERE id = ?", id)
		if err != nil {
			tx.Rollback()
			log.Printf("fail: tx.Exec, %v\n", err)
			return err
		}

		_, err = tx.Exec("UPDATE book SET likes = likes - 1 WHERE id = ?", itemId)
		if err != nil {
			tx.Rollback()
			log.Printf("fail: tx.Exec, %v\n", err)
			return err
		}
	} else if category == "video" {
		var exists bool
		err = tx.QueryRow("SELECT EXISTS (SELECT 1 FROM likes_video WHERE id = ?)", id).Scan(&exists)
		if err != nil {
			tx.Rollback()
			log.Printf("fail: tx.QueryRow, %v\n", err)
			return err
		}
		if !exists {
			tx.Rollback()
			return errors.New("Couldn't find the data of specified id")
		}

		_, err = tx.Exec("DELETE FROM likes_video WHERE id = ?", id)
		if err != nil {
			tx.Rollback()
			log.Printf("fail: tx.Exec, %v\n", err)
			return err
		}

		_, err = tx.Exec("UPDATE video SET likes = likes - 1 WHERE id = ?", itemId)
		if err != nil {
			tx.Rollback()
			log.Printf("fail: tx.Exec, %v\n", err)
			return err
		}
	} else if category == "work" {
		var exists bool
		err = tx.QueryRow("SELECT EXISTS (SELECT 1 FROM likes_work WHERE id = ?)", id).Scan(&exists)
		if err != nil {
			tx.Rollback()
			log.Printf("fail: tx.QueryRow, %v\n", err)
			return err
		}
		if !exists {
			tx.Rollback()
			return errors.New("Couldn't find the data of specified id")
		}

		_, err = tx.Exec("DELETE FROM likes_work WHERE id = ?", id)
		if err != nil {
			tx.Rollback()
			log.Printf("fail: tx.Exec, %v\n", err)
			return err
		}

		_, err = tx.Exec("UPDATE work SET likes = likes - 1 WHERE id = ?", itemId)
		if err != nil {
			tx.Rollback()
			log.Printf("fail: tx.Exec, %v\n", err)
			return err
		}
	} else {
		return errors.New("Item category is invalid")
	}

	// トランザクションを終了
	if err := tx.Commit(); err != nil {
		log.Printf("fail: tx.Commit, %v\n", err)
		return err
	}

	return nil
}

// 検索キーからブログを最終更新日時の降順で取得
//func BlogsSearch(key string) ([]model.Blog, error) {
//	query := "SELECT id, title, birth_time, update_time FROM blog WHERE (title LIKE ? OR content LIKE ?) AND public = 1 ORDER BY update_time DESC"
//	keyword := "%" + key + "%" // キーワードを含む部分文字列を作成
//	rows, err := Db.Query(query, keyword, keyword)
//	if err != nil {
//		log.Printf("fail: Db.Query, %v\n", err)
//		return nil, err
//	}
//
//	// rowsはここで、しっかり処理して、Go特有の型に変換してからusecaseに戻す
//	blogs := make([]model.Blog, 0)
//	for rows.Next() {
//		var b model.Blog
//		// ブログID、タイトル、最終更新日時、の3つだけを返す！
//		if err := rows.Scan(&b.Id, &b.Title, &b.BirthTime, &b.UpdateTime); err != nil {
//			log.Printf("fail: rows.Scan, %v\n", err)
//			if err := rows.Close(); err != nil { // 500を返して終了するが、その前にrowsのClose処理が必要
//				log.Printf("fail: rows.Close(), %v\n", err)
//			}
//			return nil, err
//		}
//		blogs = append(blogs, b) // users にデータを順々に格納！
//	}
//
//	return blogs, nil
//}

// 全てのブログを取得
//func BlogsGetAll() ([]model.Blog, error) {
//	rows, err := Db.Query("SELECT id, title, birth_time, update_time FROM blog WHERE public = 1")
//	if err != nil {
//		log.Printf("fail: Db.Query, %v\n", err)
//		return nil, err
//	}
//
//	blogs := make([]model.Blog, 0)
//	for rows.Next() {
//		var b model.Blog
//		if err := rows.Scan(&b.Id, &b.Title, &b.BirthTime, &b.UpdateTime); err != nil {
//			log.Printf("fail: rows.Scan, %v\n", err)
//			if err := rows.Close(); err != nil {
//				log.Printf("fail: rows.Close(), %v\n", err)
//			}
//			return nil, err
//		}
//		blogs = append(blogs, b)
//	}
//
//	return blogs, nil
//}

// userNameをもとにDBからデータを取得して、usecaseに返す
func LikeCheck(category string, itemId string, userId string) (model.Like, error) {
	var rows *sql.Rows // 外で変数を宣言
	// ここではトランザクションが不要！
	if category == "blog" {
		rows, err := Db.Query("SELECT * FROM likes_blog WHERE (item_id = ? AND user_id = ?)", itemId, userId)
		if err != nil {
			log.Printf("fail: Db.Query, %v\n", err)
			return model.Like{}, err
		}
		defer rows.Close()
	} else if category == "book" {
		rows, err := Db.Query("SELECT * FROM likes_book WHERE (item_id = ? AND user_id = ?)", itemId, userId)
		if err != nil {
			log.Printf("fail: Db.Query, %v\n", err)
			return model.Like{}, err
		}
		defer rows.Close()
	} else if category == "video" {
		rows, err := Db.Query("SELECT * FROM likes_video WHERE (item_id = ? AND user_id = ?)", itemId, userId)
		if err != nil {
			log.Printf("fail: Db.Query, %v\n", err)
			return model.Like{}, err
		}
		defer rows.Close()
	} else if category == "work" {
		rows, err := Db.Query("SELECT * FROM likes_work WHERE (item_id = ? AND user_id = ?)", itemId, userId)
		if err != nil {
			log.Printf("fail: Db.Query, %v\n", err)
			return model.Like{}, err
		}
		defer rows.Close()
	} else {
		return model.Like{}, errors.New("Item category is invalid")
	}

	// MySQL特有の型は、dao以外では登場させたくない → この処理はdaoで済ませる
	var l model.Like
	if rows.Next() {
		if err := rows.Scan(&l.Id, &l.UserId, &l.ItemId, &l.BirthTime); err != nil {
			log.Printf("fail: rows.Scan, %v\n", err)
			return model.Like{}, err
		}
	} else {
		// クエリ結果が空の場合
		// この場合、エラーではなく、正常値として処理したいので、nilを返す！
		return model.Like{}, nil
	}

	return l, nil
}
