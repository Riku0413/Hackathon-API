package dao

import (
	"github.com/Riku0413/Hackathon-API/model"
	"log"
)

// nameの値をもとにDBからデータを取得して、スライスの形でusecaseに返す
func GetRobotsByName(name string) ([]model.Robot, error) {
	// ここではトランザクションが不要！
	rows, err := Db.Query("SELECT id, name, age FROM robot WHERE name = ?", name)
	if err != nil {
		log.Printf("fail: Db.Query, %v\n", err)
		return nil, err
	}

	// ここで、MySQL特有の型を、model.User型に変換する
	// MySQL特有の型は、dao以外では登場させたくない → この処理はdaoで済ませる
	robots := make([]model.Robot, 0)
	for rows.Next() {
		var r model.Robot
		if err := rows.Scan(&r.Id, &r.Name, &r.Age); err != nil {
			log.Printf("fail: rows.Scan, %v\n", err)
			return nil, err
		}
		robots = append(robots, r)
	}

	return robots, nil
}

// トランザクションで新しいユーザーデータをポスト
func RegisterRobot(robot model.Robot) error {

	// ポストするデータが揃ったので、トランザクションを開始
	tx, err := Db.Begin()
	if err != nil {
		log.Printf("fail: Db.begin, %v\n", err)
		return err
	}

	// MySQLの操作
	_, err = tx.Exec("INSERT INTO robot (id, name, age) VALUES (?, ?, ?)", robot.Id, robot.Name, robot.Age)
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

func GetAllRobots() ([]model.Robot, error) {
	rows, err := Db.Query("SELECT id, name, age FROM robot")
	if err != nil {
		log.Printf("fail: Db.Query, %v\n", err)
		return nil, err
	}

	// rowsはここで、しっかり処理して、Go特有の型に変換してからusecaseに戻す
	robots := make([]model.Robot, 0)
	for rows.Next() {
		var r model.Robot
		if err := rows.Scan(&r.Id, &r.Name, &r.Age); err != nil {
			log.Printf("fail: rows.Scan, %v\n", err)
			if err := rows.Close(); err != nil { // 500を返して終了するが、その前にrowsのClose処理が必要
				log.Printf("fail: rows.Close(), %v\n", err)
			}
			return nil, err
		}
		robots = append(robots, r) // users にデータを順々に格納！
	}

	return robots, nil
}
