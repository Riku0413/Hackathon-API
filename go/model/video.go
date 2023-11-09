package model

import (
	"errors"
	"net/url"
)

type Video struct {
	Id           string   `json:"id"`
	Title        string   `json:"title"`
	URL          *url.URL `json:"url"`
	UserId       string   `json:"user_id"`
	BirthTime    string   `json:"birth_time"`
	UpdateTime   string   `json:"update_time"`
	Public       bool     `json:"publish"`
	Introduction string   `json:"introduction"`
}

func (v Video) VideoCheckUID() error {
	if len(v.Id) != 26 {
		return errors.New("The UID format is invalid")
	}
	return nil
}
