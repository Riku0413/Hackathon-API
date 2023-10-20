package model

import "errors"

// DBから取得するデータの枠組みを設定
type User struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// Nameのバリデーションメソッド
func (u User) CheckName() error {
	if u.Name == "" {
		return errors.New("名前は空欄にできません")
	} else if len(u.Name) > 50 {
		return errors.New("名前は50文字以下である必要があります")
	}
	return nil
}

func (u User) CheckAge() error {
	if u.Age < 20 {
		return errors.New("20歳未満のユーザーは登録できません")
	} else if len(u.Name) > 80 {
		return errors.New("80歳以下ののユーザーしか登録できません")
	}
	return nil
}
