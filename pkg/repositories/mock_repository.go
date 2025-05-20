package repositories

import (
	"errors"

	"github.com/Finanzas-3438/InverBono.git/pkg/domain/entities"
	"github.com/Finanzas-3438/InverBono.git/pkg/interfaces"
)

type InMemoryUserRepo struct {
	users []*entities.User
}

func NewInMemoryUserRepo() interfaces.UserRepository {
	return &InMemoryUserRepo{
		users: []*entities.User{
			&entities.User{
				ID:       1,
				Username: "admin",
				Password: "1234",
				Role:     "admin",
			},
		},
	}
}

func (r *InMemoryUserRepo) GetByUsername(username string) (*entities.User, error) {
	for _, user := range r.users {
		if user.Username == username {
			return user, nil
		}
	}
	return nil, errors.New("usuario no encontrado")
}

func (r *InMemoryUserRepo) SaveUser(user *entities.User) error {
	for _, u := range r.users {
		if u.Username == user.Username {
			return errors.New("usuario ya existe")
		}
	}
	r.users = append(r.users, user)
	return nil
}

type InMemoryProfileRepo struct {
	profiles []*entities.Profile
}

func NewInMemoryProfileRepo() interfaces.ProfileRepository {
	fakeProfile := entities.NewFakeProfile()
	profiles := []*entities.Profile{
		&fakeProfile,
	}
	return &InMemoryProfileRepo{
		profiles: profiles,
	}
}
func (r *InMemoryProfileRepo) GetByID(id int) (*entities.Profile, error) {
	for _, profile := range r.profiles {
		if profile.ID == id {
			return profile, nil
		}
	}
	return nil, errors.New("perfil no encontrado")
}
func (r *InMemoryProfileRepo) SaveProfile(profile entities.Profile) error {
	r.profiles = append(r.profiles, &profile)
	return nil
}
func (r *InMemoryProfileRepo) UpdateProfile(profile entities.Profile) error {
	for i, p := range r.profiles {
		if p.ID == profile.ID {
			r.profiles[i] = &profile
			return nil
		}
	}
	return errors.New("perfil no encontrado")
}
