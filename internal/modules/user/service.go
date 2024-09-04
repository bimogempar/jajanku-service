package user

import (
	"errors"
	"jajanku_service/domain"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct {
	UserRepository domain.UserRepository
	JWTSecret      string
}

func NewUserService(repo domain.UserRepository, secret string) domain.UserService {
	return &UserServiceImpl{
		UserRepository: repo,
		JWTSecret:      secret,
	}
}

func (s *UserServiceImpl) RegisterUser(user *domain.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return s.UserRepository.CreateUser(user)
}

func (s *UserServiceImpl) LoginUser(email, password string) (string, error) {
	user, err := s.UserRepository.GetUserByEmail(email)
	if err != nil {
		return "", errors.New("user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid credentials")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte(s.JWTSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *UserServiceImpl) GetAllUsers() ([]*domain.User, error) {
	users, err := s.UserRepository.GetAllUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}
