package user

import "github.com/jmoiron/sqlx"

type UserRepository struct {
	DB *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

// creates a new user
func (r *UserRepository) createUser(signUpReq SignUpReq) error {
	request := `INSERT INTO users(name, email, password) VALUES (:name, :email, :password)`

	_, err := r.DB.NamedExec(request, signUpReq)

	if err != nil {
		return err
	}

	return nil
}
