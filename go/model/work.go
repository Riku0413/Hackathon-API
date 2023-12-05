package model

import (
	"errors"
)

type Work struct {
	Id           string `json:"id"`
	Title        string `json:"title"`
	URL          string `json:"url"`
	UserId       string `json:"user_id"`
	BirthTime    string `json:"birth_time"`
	UpdateTime   string `json:"update_time"`
	Public       bool   `json:"publish"`
	Introduction string `json:"introduction"`
}

func (w Work) WorkCheckUID() error {
	if len(w.Id) != 26 {
		return errors.New("The UID format is invalid")
	}
	return nil
}
