package usecases

import (
	"time"

	"github.com/Finanzas-3438/InverBono.git/pkg/interfaces"
	"github.com/golang-jwt/jwt/v5"
)

type AuthUseCase struct {
	UserRepo interfaces.UserRepository
	Secret   string
}

func NewAuthUseCase(repo interfaces.UserRepository, secret string) *AuthUseCase {
	return &AuthUseCase{
		UserRepo: repo,
		Secret:   secret,
	}
}

func (a *AuthUseCase) Login(username, password string) (string, error) {
	user, err := a.UserRepo.GetByUsername(username)
	if err != nil || user.Password != password {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"role":     user.Role,
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
	})

	tokenString, err := token.SignedString([]byte(a.Secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
