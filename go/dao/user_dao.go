package dao

import (
	"errors"
	"github.com/Riku0413/Hackathon-API/model"
	"log"
	"time"
)

// nameの値をもとにDBからデータを取得して、スライスの形でusecaseに返す
//func GetUsersByName(name string) ([]model.User, error) {
//	// ここではトランザクションが不要！
//	rows, err := Db.Query("SELECT id, name, age FROM user_ver_0 WHERE name = ?", name)
//	if err != nil {
//		log.Printf("fail: Db.Query, %v\n", err)
//		return nil, err
//	}
//
//	// ここで、MySQL特有の型を、model.User型に変換する
//	// MySQL特有の型は、dao以外では登場させたくない → この処理はdaoで済ませる
//	users := make([]model.User, 0)
//	for rows.Next() {
//		var u model.User
//		if err := rows.Scan(&u.Id, &u.Name, &u.Age); err != nil {
//			log.Printf("fail: rows.Scan, %v\n", err)
//			return nil, err
//		}
//		users = append(users, u)
//	}
//
//	return users, nil
//}

// トランザクションで新しいユーザーデータをポスト
func UserSignUp(user model.User) error {

	// ポストするデータが揃ったので、トランザクションを開始
	tx, err := Db.Begin()
	if err != nil {
		log.Printf("fail: Db.begin, %v\n", err)
		return err
	}

	// 1. 指定したuidがデータベースに存在するか確認
	// このバリデーションはデータベースと接続して確認する必要があるからdaoで。
	var count int
	err = tx.QueryRow("SELECT COUNT(*) FROM user WHERE id = ?", user.Id).Scan(&count)
	if err != nil {
		tx.Rollback()
		log.Printf("fail: checking existing uid, %v\n", err)
		return err
	}

	// 2. uidが既に存在する場合はエラーを返す
	if count > 0 {
		tx.Rollback()
		log.Println("fail: UID already exists")
		return errors.New("UID already exists")
	}

	// 3. MySQLの操作
	_, err = tx.Exec("INSERT INTO user (id, first_name, family_name, register_time, last_online_time) VALUES (?, ?, ?, ?, ?)", user.Id, nil, nil, time.Now(), time.Now())
	if err != nil {
		// エラーが発生した場合、トランザクションをロールバックし、エラーを上位に返す
		tx.Rollback()
		log.Printf("fail: tx.Exec, %v\n", err)
		// log.Fatal(err) // これいるの？　→ データベースがバグるから不要！
		return err
	}

	// トランザクションを終了
	if err := tx.Commit(); err != nil {
		log.Printf("fail: tx.Commit, %v\n", err)
		return err
	}

	return nil
}

//func GetAllUsers() ([]model.User, error) {
//	rows, err := Db.Query("SELECT id, name, age FROM user_ver_0")
//	if err != nil {
//		log.Printf("fail: Db.Query, %v\n", err)
//		return nil, err
//	}

// rowsはここで、しっかり処理して、Go特有の型に変換してからusecaseに戻す
//users := make([]model.User, 0)
//for rows.Next() {
//	var u model.User
//	if err := rows.Scan(&u.Id, &u.Name, &u.Age); err != nil {
//		log.Printf("fail: rows.Scan, %v\n", err)
//		if err := rows.Close(); err != nil { // 500を返して終了するが、その前にrowsのClose処理が必要
//			log.Printf("fail: rows.Close(), %v\n", err)
//		}
//		return nil, err
//	}
//	users = append(users, u) // users にデータを順々に格納！
//}

//return users, nil
//}
