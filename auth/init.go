package auth

import (
	"github.com/erdedigital/go-users/auth/model"
	"github.com/erdedigital/go-users/auth/service"
)

func CreateAuth(p *model.Auth, s service.Auth) error {
	err := s.Save(p)
	if err != nil {
		return err
	}
	return nil
}
