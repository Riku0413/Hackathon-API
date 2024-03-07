package dao

import (
	"errors"
	"github.com/Riku0413/Hackathon-API/model"
	"log"
)

// トランザクションで新しいユーザーデータをポスト
func BookPost(book model.Book) error {

	// ポストするデータが揃ったので、トランザクションを開始
	tx, err := Db.Begin()
	if err != nil {
		log.Printf("fail: Db.begin, %v\n", err)
		return err
	}

	// MySQLの操作
	_, err = tx.Exec("INSERT INTO book (id, title, user_id, birth_time, update_time, public) VALUES (?, ?, ?, ?, ?, ?)", book.Id, book.Title, book.UserId, book.BirthTime, book.UpdateTime, book.Public)
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
func BookPut(book model.Book) error {

	// ポストするデータが揃ったので、トランザクションを開始
	tx, err := Db.Begin()
	if err != nil {
		log.Printf("fail: Db.begin, %v\n", err)
		return err
	}

	// MySQLの操作
	_, err = tx.Exec("UPDATE book SET title = ?, introduction = ?, update_time = ? WHERE id = ?", book.Title, book.Introduction, book.UpdateTime, book.Id)
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
func BookPublish(book model.Book) error {

	// ポストするデータが揃ったので、トランザクションを開始
	tx, err := Db.Begin()
	if err != nil {
		log.Printf("fail: Db.begin, %v\n", err)
		return err
	}

	// MySQLの操作
	_, err = tx.Exec("UPDATE book SET update_time = ?, public = ? WHERE id = ?", book.UpdateTime, book.Public, book.Id)
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
func BookGet(id string) (model.Book, error) {
	// ここではトランザクションが不要！
	rows, err := Db.Query("SELECT * FROM book WHERE id = ?", id)
	if err != nil {
		log.Printf("fail: Db.Query, %v\n", err)
		return model.Book{}, err
	}
	defer rows.Close()

	// MySQL特有の型は、dao以外では登場させたくない → この処理はdaoで済ませる
	var b model.Book
	if rows.Next() {
		// ここのコラムの順番はデータベースに合わせる必要がある！！！
		if err := rows.Scan(&b.Id, &b.UserId, &b.Title, &b.Introduction, &b.BirthTime, &b.UpdateTime, &b.Public); err != nil {
			log.Printf("fail: rows.Scan, %v\n", err)
			return model.Book{}, err
		}
	} else {
		// クエリ結果が空の場合
		return model.Book{}, errors.New("Couldn't find the book specified")
	}

	return b, nil
}

// idをもとにDBからブログを削除して、削除したブログのidをusecaseに返す
func BookDelete(id string) error {

	// 削除するブログのidが決まったので、トランザクションを開始
	tx, err := Db.Begin()
	if err != nil {
		log.Printf("fail: Db.begin, %v\n", err)
		return err
	}

	// データの存在性の確認
	var exists bool
	err = tx.QueryRow("SELECT EXISTS (SELECT 1 FROM book WHERE id = ?)", id).Scan(&exists)
	if err != nil {
		tx.Rollback()
		log.Printf("fail: tx.QueryRow, %v\n", err)
		return err
	}
	if !exists {
		tx.Rollback()
		return errors.New("Couldn't find the data of specified id")
	}

	// 本のページを削除
	_, err = tx.Exec("DELETE FROM chapter WHERE book_id = ?", id)
	if err != nil {
		tx.Rollback()
		log.Printf("fail: tx.Exec, %v\n", err)
		return err
	}

	// 本を削除
	_, err = tx.Exec("DELETE FROM book WHERE id = ?", id)
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
func BooksGet(userId string) ([]model.Book, error) {
	rows, err := Db.Query("SELECT id, title, introduction, update_time FROM book WHERE user_id = ? ORDER BY update_time DESC", userId)
	if err != nil {
		log.Printf("fail: Db.Query, %v\n", err)
		return nil, err
	}

	// rowsはここで、しっかり処理して、Go特有の型に変換してからusecaseに戻す
	books := make([]model.Book, 0)
	for rows.Next() {
		var b model.Book
		// ブログID、タイトル、最終更新日時、の3つだけを返す！
		if err := rows.Scan(&b.Id, &b.Title, &b.Introduction, &b.UpdateTime); err != nil {
			log.Printf("fail: rows.Scan, %v\n", err)
			if err := rows.Close(); err != nil { // 500を返して終了するが、その前にrowsのClose処理が必要
				log.Printf("fail: rows.Close(), %v\n", err)
			}
			return nil, err
		}
		books = append(books, b) // users にデータを順々に格納！
	}

	return books, nil
}

// 検索キーから本を最終更新日時の降順で取得
func BooksSearch(key string) ([]model.Book, error) {
	query := "SELECT id, title, birth_time, update_time FROM book WHERE (title LIKE ?) AND public = 1 ORDER BY update_time DESC"
	keyword := "%" + key + "%" // キーワードを含む部分文字列を作成
	rows, err := Db.Query(query, keyword)
	if err != nil {
		log.Printf("fail: Db.Query, %v\n", err)
		return nil, err
	}

	// rowsはここで、しっかり処理して、Go特有の型に変換してからusecaseに戻す
	books := make([]model.Book, 0)
	for rows.Next() {
		var b model.Book
		// 本ID、タイトル、最終更新日時、の3つだけを返す！
		if err := rows.Scan(&b.Id, &b.Title, &b.BirthTime, &b.UpdateTime); err != nil {
			log.Printf("fail: rows.Scan, %v\n", err)
			if err := rows.Close(); err != nil { // 500を返して終了するが、その前にrowsのClose処理が必要
				log.Printf("fail: rows.Close(), %v\n", err)
			}
			return nil, err
		}
		books = append(books, b) // users にデータを順々に格納！
	}

	return books, nil
}

// 全ての本を取得
func BooksGetAll() ([]model.Book, error) {
	rows, err := Db.Query("SELECT id, title, birth_time, update_time FROM book WHERE public = 1")
	if err != nil {
		log.Printf("fail: Db.Query, %v\n", err)
		return nil, err
	}

	books := make([]model.Book, 0)
	for rows.Next() {
		var b model.Book
		if err := rows.Scan(&b.Id, &b.Title, &b.BirthTime, &b.UpdateTime); err != nil {
			log.Printf("fail: rows.Scan, %v\n", err)
			if err := rows.Close(); err != nil { // 500を返して終了するが、その前にrowsのClose処理が必要
				log.Printf("fail: rows.Close(), %v\n", err)
			}
			return nil, err
		}
		books = append(books, b)
	}

	return books, nil
}
