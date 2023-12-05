package model

import "errors"

type Like struct {
	Id           string `json:"id"`
	UserId       string `json:"user_id"`
	ItemCategory string `json:"item_category"`
	ItemId       string `json:"item_id"`
	BirthTime    string `json:"birth_time"`
}

func (l Like) LikeCheckUID() error {
	if len(l.Id) != 26 {
		return errors.New("UIDが不正なフォーマットです")
	}
	return nil
}
