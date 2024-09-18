package user

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Login(email, password string) (string, error)
	Register(user *User) (*User, error)
	GetUser(id string) (*User, error)
	ListUsers() ([]*User, error)
}

type UserService struct {
	repo      Repository
	jwtSecret string
}

func NewUserService(repo Repository, jwtSecret string) Service {
	return &UserService{
		repo:      repo,
		jwtSecret: jwtSecret,
	}
}

func (s *UserService) Login(email, password string) (string, error) {
	user, err := s.repo.GetByEmail(email)
	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid credentials")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte(s.jwtSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *UserService) Register(user *User) (*User, error) {
	if user.Role == "" {
		user.Role = "user"
	}

	existUser, err := s.repo.GetByEmail(user.Email)
	if err != nil {
		return nil, err
	}
	if existUser != nil {
		return nil, fmt.Errorf("email %s already in use", existUser.Email)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.Password = string(hashedPassword)
	return user, s.repo.Create(user)
}

func (s *UserService) GetUser(id string) (*User, error) {
	return s.repo.GetByID(id)
}

func (s *UserService) ListUsers() ([]*User, error) {
	return s.repo.GetAll()
}
