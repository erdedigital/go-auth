package service

import (
	"database/sql"

	"github.com/erdedigital/go-users/auth/model"
)

type Auth interface {
	Save(*model.Auth) error
}

type auth struct {
	db *sql.DB
}

func NewAuth(db *sql.DB) *auth {
	return &auth{db}
}

func (p *auth) Save(u *model.Auth) error {
	query := `INSERT INTO public.users(email) VALUES($1)`
	statement, err := p.db.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(u.Email)

	if err != nil {
		return err
	}

	return nil
}
