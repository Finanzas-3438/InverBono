package interfaces

import "github.com/Finanzas-3438/InverBono.git/pkg/domain/entities"

type UserRepository interface {
	GetByUsername(username string) (*entities.User, error)
	SaveUser(user *entities.User) error
}

type ProfileRepository interface {
	GetByID(id int) (*entities.Profile, error)
	SaveProfile(profile entities.Profile) error
	UpdateProfile(profile entities.Profile) error
}
