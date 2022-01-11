package auth

import (
	"project/playground/be5/docker-play/entities"

	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *AuthRepository {
	return &AuthRepository{
		db: db,
	}
}

func (a *AuthRepository) Login(name, hp string) (entities.User, error) {
	loggedInUser := entities.User{Nama: name, HP: hp}

	if err := a.db.First(&loggedInUser).Error; err != nil {
		return loggedInUser, err
	}
	return loggedInUser, nil
}
