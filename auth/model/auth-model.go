package model

import "time"

type Auth struct {
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
	Email     string `json:"email"`
}

func NewAuth() *Auth {
	return &Auth{
		CreatedAt: DateNow(),
		UpdatedAt: DateNow(),
	}
}

func DateNow() int64 {
	return time.Now().Unix()
}
