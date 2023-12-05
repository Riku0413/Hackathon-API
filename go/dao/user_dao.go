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
	_, err = tx.Exec("INSERT INTO user (id, user_name, introduction, git_hub, register_time, last_online_time) VALUES (?, ?, ?, ?, ?, ?)", user.Id, user.UserName, "", "", time.Now(), time.Now())
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

// userNameをもとにDBからデータを取得して、usecaseに返す
func UserCheck(userName string) (model.User, error) {
	// ここではトランザクションが不要！
	rows, err := Db.Query("SELECT * FROM user WHERE user_name = ?", userName)
	if err != nil {
		log.Printf("fail: Db.Query, %v\n", err)
		return model.User{}, err
	}
	defer rows.Close()

	// MySQL特有の型は、dao以外では登場させたくない → この処理はdaoで済ませる
	var u model.User
	if rows.Next() {
		if err := rows.Scan(&u.Id, &u.RegisterTime, &u.LastTime, &u.UserName, &u.Introduction, &u.GitHub); err != nil {
			log.Printf("fail: rows.Scan, %v\n", err)
			return model.User{}, err
		}
	} else {
		// クエリ結果が空の場合
		// この場合、エラーではなく、正常値として処理したいので、nilを返す！
		return model.User{}, nil
	}

	return u, nil
}

// userNameをもとにDBからデータを取得して、usecaseに返す
func UserGet(id string) (model.User, error) {
	// ここではトランザクションが不要！
	rows, err := Db.Query("SELECT * FROM user WHERE id = ?", id)
	if err != nil {
		log.Printf("fail: Db.Query, %v\n", err)
		return model.User{}, err
	}
	defer rows.Close()

	// MySQL特有の型は、dao以外では登場させたくない → この処理はdaoで済ませる
	var u model.User
	if rows.Next() {
		if err := rows.Scan(&u.Id, &u.RegisterTime, &u.LastTime, &u.UserName, &u.Introduction, &u.GitHub); err != nil {
			log.Printf("fail: rows.Scan, %v\n", err)
			return model.User{}, err
		}
	} else {
		// クエリ結果が空の場合
		return model.User{}, errors.New("Couldn't find the user specified")
	}

	return u, nil
}

// ユーザープロフィールの更新
func UserUpdate(user model.User) error {
	tx, err := Db.Begin()
	if err != nil {
		log.Printf("fail: Db.begin, %v\n", err)
		return err
	}

	// MySQLの操作
	_, err = tx.Exec("UPDATE user SET user_name = ?, introduction = ?, git_hub = ? WHERE id = ?", user.UserName, user.Introduction, user.GitHub, user.Id)
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
