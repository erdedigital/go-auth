package auth

import "database/sql"

type AuthService interface {
	ServiceSave(*ModuleUser) error
}

type authService struct {
	db *sql.DB
}

func NewAuthService(db *sql.DB) *authService {
	return &authService{db}
}

func (p *authService) ServiceSave(u *ModuleUser) error {
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
