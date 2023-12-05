package dao

import (
	"errors"
	"github.com/Riku0413/Hackathon-API/model"
	"log"
)

// トランザクションで新しいブログをポスト
func ChapterPost(chapter model.Chapter) error {

	// ポストするデータが揃ったので、トランザクションを開始
	tx, err := Db.Begin()
	if err != nil {
		log.Printf("fail: Db.begin, %v\n", err)
		return err
	}

	// MySQLの操作
	_, err = tx.Exec("INSERT INTO chapter (id, book_id, title, content, birth_time, update_time) VALUES (?, ?, ?, ?, ?, ?)", chapter.Id, chapter.BookId, chapter.Title, chapter.Content, chapter.BirthTime, chapter.UpdateTime)
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
func ChapterPut(chapter model.Chapter) error {

	// ポストするデータが揃ったので、トランザクションを開始
	tx, err := Db.Begin()
	if err != nil {
		log.Printf("fail: Db.begin, %v\n", err)
		return err
	}

	// MySQLの操作
	_, err = tx.Exec("UPDATE chapter SET title = ?, content = ?, update_time = ? WHERE id = ?", chapter.Title, chapter.Content, chapter.UpdateTime, chapter.Id)
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
func ChapterGet(id string) (model.Chapter, error) {
	// ここではトランザクションが不要！
	rows, err := Db.Query("SELECT * FROM chapter WHERE id = ?", id)
	if err != nil {
		log.Printf("fail: Db.Query, %v\n", err)
		return model.Chapter{}, err
	}
	defer rows.Close()

	// MySQL特有の型は、dao以外では登場させたくない → この処理はdaoで済ませる
	var c model.Chapter
	if rows.Next() {
		if err := rows.Scan(&c.Id, &c.BookId, &c.Title, &c.Content, &c.BirthTime, &c.UpdateTime); err != nil {
			log.Printf("fail: rows.Scan, %v\n", err)
			return model.Chapter{}, err
		}
	} else {
		// クエリ結果が空の場合
		return model.Chapter{}, errors.New("Couldn't find the chapter specified")
	}

	return c, nil
}

// user_idに紐づくブログを最終更新日時の降順で取得
func ChaptersGet(bookId string) ([]model.Chapter, error) {
	rows, err := Db.Query("SELECT id, title, update_time FROM chapter WHERE book_id = ? ORDER BY birth_time ASC", bookId)
	if err != nil {
		log.Printf("fail: Db.Query, %v\n", err)
		return nil, err
	}

	// rowsはここで、しっかり処理して、Go特有の型に変換してからusecaseに戻す
	chapters := make([]model.Chapter, 0)
	for rows.Next() {
		var c model.Chapter
		// チャプターID、タイトル、最終更新日時、の3つだけを返す！
		if err := rows.Scan(&c.Id, &c.Title, &c.UpdateTime); err != nil {
			log.Printf("fail: rows.Scan, %v\n", err)
			if err := rows.Close(); err != nil { // 500を返して終了するが、その前にrowsのClose処理が必要
				log.Printf("fail: rows.Close(), %v\n", err)
			}
			return nil, err
		}
		chapters = append(chapters, c) // users にデータを順々に格納！
	}

	return chapters, nil
}
