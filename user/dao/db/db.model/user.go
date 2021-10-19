package db_model

import "time"

type Users struct {
	Id           int64     `json:"id"`
	Name         string    `json:"name"`
	Status       int8      `json:"status"`
	Avatar       string    `json:"avatar"`
	NickName     string    `json:"nick_name"`
	SignInfo     string    `json:"sign_info"`
	Gender       int8      `json:"gender"`
	Birthday     string    `json:"birthday"`
	Account      string    `json:"account"`
	Email        string    `json:"email"`
	CountryCode  string    `json:"country_code"`
	Mobile       string    `json:"mobile"`
	RegisterType int8      `json:"register_type"`
	Password     string    `json:"password"`
	Salt         string    `json:"salt"`
	CreateTime   time.Time `json:"create_time"`
	UpdateTime   time.Time `json:"update_time"`
}
