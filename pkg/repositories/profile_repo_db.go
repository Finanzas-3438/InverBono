package repositories

import (
	"github.com/Finanzas-3438/InverBono.git/pkg/domain/entities"
	"gorm.io/gorm"
)

type ProfileRepoDB struct {
	DB *gorm.DB
}

func (r *ProfileRepoDB) GetByID(id int) (*entities.Profile, error) {
	var profile entities.Profile
	err := r.DB.Where("id = ?", id).First(&profile).Error
	if err != nil {
		return nil, err
	}
	return &profile, nil
}

func (r *ProfileRepoDB) SaveProfile(profile entities.Profile) error {
	return r.DB.Create(&profile).Error
}
func (r *ProfileRepoDB) UpdateProfile(profile entities.Profile) error {
	return r.DB.Save(&profile).Error
}
