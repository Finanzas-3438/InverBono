package usecases

import (
	"errors"
	"log"

	"github.com/Finanzas-3438/InverBono.git/pkg/domain/entities"
	"github.com/Finanzas-3438/InverBono.git/pkg/interfaces"
)

type ProfileUsecase interface {
	GetProfileByID(id int, actor *entities.User) (*entities.Profile, error)
	SaveProfile(profile entities.Profile, actor *entities.User) error
	UpdateProfile(profile entities.Profile, actor *entities.User) error
}

type profileUC struct {
	ProfileRepo interfaces.ProfileRepository
	UserRepo    interfaces.UserRepository
}

func NewProfileUseCase(repo interfaces.ProfileRepository, userRepo interfaces.UserRepository) ProfileUsecase {
	return &profileUC{
		ProfileRepo: repo,
		UserRepo:    userRepo,
	}
}
func (p *profileUC) GetProfileByID(id int, actor *entities.User) (*entities.Profile, error) {

	log.Printf("GetProfileByID: %d", id)
	profile, err := p.ProfileRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return profile, nil
}

func (p *profileUC) SaveProfile(profile entities.Profile, actor *entities.User) error {

	if profile.UserID != int(actor.ID) {
		return errors.New("no tienes permiso para guardar este perfil")
	}

	err := p.ProfileRepo.SaveProfile(profile)
	if err != nil {
		return err
	}

	return nil
}
func (p *profileUC) UpdateProfile(profile entities.Profile, actor *entities.User) error {
	err := p.ProfileRepo.UpdateProfile(profile)
	if err != nil {
		return err
	}

	return nil
}
