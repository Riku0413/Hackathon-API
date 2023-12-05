package model

import "errors"

// DBから取得するデータの枠組みを設定
type User struct {
	Id           string `json:"id"`
	UserName     string `json:"user_name"`
	Introduction string `json:"introduction"`
	GitHub       string `json:"git_hub"`
	RegisterTime string `json:"register_time"`
	LastTime     string `json:"last_time"`
}

// Nameのバリデーションメソッド
func (u User) UserCheckUID() error {
	if len(u.Id) != 28 {
		return errors.New("The UID has an invalid format")
	}
	return nil
}
