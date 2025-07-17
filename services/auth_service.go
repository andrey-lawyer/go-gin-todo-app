package services

import (
	"github.com/andrey-lawyer/go-gin-todo-app/errors"
	"github.com/andrey-lawyer/go-gin-todo-app/models"
	"github.com/andrey-lawyer/go-gin-todo-app/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserRepo *repositories.UserRepository
}

func NewAuthService(userRepo *repositories.UserRepository) *AuthService {
	return &AuthService{UserRepo: userRepo}
}

func (s *AuthService) RegisterUser(username, password string) (*models.User, error) {
	if err := errors.ValidatePassword(password); err != nil {
		return nil, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		ID:       primitive.NewObjectID(),
		Username: username,
		Password: string(hashedPassword),
	}

	err = s.UserRepo.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *AuthService) LoginUser(username, password string) (*models.User, error) {
	user, err := s.UserRepo.FindByUsername(username)
	if err != nil {
		return nil, &errors.AuthError{Msg: "invalid credentials"}
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		// Возвращаем бизнес-ошибку, если пароль неверный
		return nil, &errors.AuthError{Msg: "invalid credentials"}
	}

	return user, nil
}
