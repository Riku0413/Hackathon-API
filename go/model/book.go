package model

import "errors"

type Book struct {
	Id           string `json:"id"`
	Title        string `json:"title"`
	Introduction string `json:"introduction"`
	UserId       string `json:"user_id"`
	BirthTime    string `json:"birth_time"`
	UpdateTime   string `json:"update_time"`
	Public       bool   `json:"publish"`
}

// Nameのバリデーションメソッド
func (b Book) BookCheckUID() error {
	if len(b.Id) != 26 {
		return errors.New("UIDが不正なフォーマットです")
	}
	return nil
}
