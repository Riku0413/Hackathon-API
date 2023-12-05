package model

import "errors"

type Chapter struct {
	Id         string `json:"id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	BookId     string `json:"book_id"`
	BirthTime  string `json:"birth_time"`
	UpdateTime string `json:"update_time"`
}

// バリデーションメソッド
func (c Chapter) ChapterCheckUID() error {
	if len(c.Id) != 26 {
		return errors.New("UIDが不正なフォーマットです")
	}
	return nil
}
