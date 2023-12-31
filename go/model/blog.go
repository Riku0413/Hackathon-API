package model

import "errors"

type Blog struct {
	Id         string `json:"id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	UserId     string `json:"user_id"`
	BirthTime  string `json:"birth_time"`
	UpdateTime string `json:"update_time"`
	Public     bool   `json:"publish"`
	Likes      int    `json:"likes"`
}

// Nameのバリデーションメソッド
func (b Blog) BlogCheckUID() error {
	if len(b.Id) != 26 {
		return errors.New("UIDが不正なフォーマットです")
	}
	return nil
}
