package repositories

import (
	"errors"

	"github.com/Finanzas-3438/InverBono.git/pkg/domain/entities"
)

type FakeUserRepo struct{}

var users = map[string]entities.User{}

func (f *FakeUserRepo) GetByUsername(username string) (*entities.User, error) {
	user, exists := users[username]
	if !exists {
		return nil, errors.New("usuario no encontrado")
	}
	return &user, nil
}

func (f *FakeUserRepo) SaveUser(user entities.User) error {
	if _, exists := users[user.Username]; exists {
		return errors.New("usuario ya existe")
	}
	users[user.Username] = user
	return nil
}
