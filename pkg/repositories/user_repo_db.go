package repositories

import (
	"github.com/Finanzas-3438/InverBono.git/pkg/domain/entities"
	"gorm.io/gorm"
)

type UserRepoDB struct {
	DB *gorm.DB
}

func (r *UserRepoDB) GetByUsername(username string) (*entities.User, error) {
	var user entities.User
	err := r.DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepoDB) SaveUser(user *entities.User) error {
	return r.DB.Create(user).Error
}
