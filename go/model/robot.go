package model

import "errors"

// DBから取得するデータの枠組みを設定
type Robot struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// Nameのバリデーションメソッド
func (r Robot) CheckName() error {
	if r.Name == "" {
		return errors.New("名前は空欄にできません")
	} else if len(r.Name) > 50 {
		return errors.New("名前は50文字以下である必要があります")
	}
	return nil
}

func (r Robot) CheckAge() error {
	if r.Age < 20 {
		return errors.New("20歳未満のロボットは登録できません")
	} else if len(r.Name) > 80 {
		return errors.New("80歳以下ののロボットしか登録できません")
	}
	return nil
}
