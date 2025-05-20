package usecases_test

import (
	"testing"

	"github.com/Finanzas-3438/InverBono.git/pkg/domain/entities"
	"github.com/Finanzas-3438/InverBono.git/pkg/domain/usecases"
	"github.com/Finanzas-3438/InverBono.git/pkg/repositories"
)

func Setup() usecases.ProfileUsecase {
	repo := repositories.NewInMemoryProfileRepo()
	userRepo := repositories.NewInMemoryUserRepo()
	uc := usecases.NewProfileUseCase(repo, userRepo)
	return uc
}
func TestSaveProfile(t *testing.T) {

	uc := Setup()

	profile := entities.NewFakeProfile()
	actor := entities.NewFakeUser()
	err := uc.SaveProfile(profile, &actor)
	if err != nil {
		t.Errorf("Error al guardar el perfil: %v", err)
	}

}
