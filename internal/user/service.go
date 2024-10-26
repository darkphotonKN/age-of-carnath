package user

import "golang.org/x/crypto/bcrypt"

type UserService struct {
	repo *UserRepository
}

func NewUserService(repo *UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}

}

func (s *UserService) signUpService(signUpReq SignUpReq) error {
	// hash password
	hashedPw, err := s.HashPassword(signUpReq.Password)

	if err != nil {
		return err
	}

	// replace password with hashed one
	signUpReq.Password = hashedPw

	// create user
	return s.repo.createUser(signUpReq)
}

// HashPassword hashes the given password using bcrypt.
func (s *UserService) HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
