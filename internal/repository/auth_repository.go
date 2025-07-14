package repository

import (
	"portfolio-backend/config"
	"portfolio-backend/internal/domain/entity"
    "gorm.io/gorm"
)

type AuthRepository struct{
    DB *gorm.DB
}

func (r *AuthRepository) FindByUsername(username string) (*entity.User, error) {
	var user entity.User
	err := config.DB.Table("users").Where("username = ?", username).First(&user).Error
	return &user, err
}

func (r *AuthRepository) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	err := config.DB.Table("users").Where("email = ?", email).First(&user).Error
	return &user, err
}

func (r *AuthRepository) Create(user *entity.User) error {
	return config.DB.Create(user).Error
}
