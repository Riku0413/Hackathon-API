package model

import "errors"

type Comment struct {
	Id           string `json:"id"`
	Content      string `json:"content"`
	UserId       string `json:"user_id"`
	UserName     string `json:"user_name"`
	ItemCategory string `json:"item_category"`
	ItemId       string `json:"item_id"`
	BirthTime    string `json:"birth_time"`
	UpdateTime   string `json:"update_time"`
}

// Nameのバリデーションメソッド
func (c Comment) CommentCheckUID() error {
	if len(c.Id) != 26 {
		return errors.New("UIDが不正なフォーマットです")
	}
	return nil
}
