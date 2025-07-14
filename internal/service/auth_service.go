package service

import (
	"errors"
	"portfolio-backend/internal/domain/entity"
	"portfolio-backend/internal/repository"
	"time"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

var ErrEmailExists = errors.New("email already exists")
var ErrInvalidLogin = errors.New("invalid username or password")

type AuthService struct {
	Repo   *repository.AuthRepository
	JwtKey []byte
}

type Claims struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

func (s *AuthService) Register(input entity.User) error {
	// Cek email
	exist, _ := s.Repo.FindByEmail(input.Email)
	if exist.ID != 0 {
		return ErrEmailExists
	}

	// Hash
	hashed, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	input.Password = string(hashed)

	return s.Repo.Create(&input)
}

func (s *AuthService) Login(username, password string) (*entity.User, string, error) {
	user, err := s.Repo.FindByUsername(username)
	if err != nil {
		return nil, "", ErrInvalidLogin
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, "", ErrInvalidLogin
	}

	exp := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exp.Unix(),
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(s.JwtKey)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}
