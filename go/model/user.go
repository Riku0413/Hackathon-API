package model

import "errors"

// DBから取得するデータの枠組みを設定
type User struct {
	Id           string `json:"id"`
	FirstName    string `json:"first_name"`
	FamilyName   string `json:"family_name"`
	RegisterTime int    `json:"register_time"`
	LastTime     int    `json:"last_time"`
}

// Nameのバリデーションメソッド
func (u User) UserCheckUID() error {
	if len(u.Id) != 28 {
		return errors.New("The UID has an invalid format")
	}
	return nil
}
