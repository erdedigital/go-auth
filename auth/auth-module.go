package auth

import "time"

type ModuleUser struct {
	CreatedAt int64  `json:"created_at,omitempty"`
	UpdatedAt int64  `json:"updated_at,omitempty"`
	Email     string `json:"email"`
}

func NewModuleUser() *ModuleUser {
	return &ModuleUser{
		CreatedAt: DateNow(),
		UpdatedAt: DateNow(),
	}
}

func DateNow() int64 {
	return time.Now().Unix()
}
