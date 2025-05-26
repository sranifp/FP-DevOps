package repository

import (
	"FP-DevOps/entity"

	"gorm.io/gorm"
)

type (
	UserRepository interface {
		Create(entity.User) (entity.User, error)
		GetUserById(string) (entity.User, error)
		GetUserByUsername(string) (entity.User, error)
	}

	userRepository struct {
		db *gorm.DB
	}
)

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) Create(user entity.User) (entity.User, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (r *userRepository) GetUserById(userID string) (entity.User, error) {
	var user entity.User
	if err := r.db.Where("id = ?", userID).Take(&user).Error; err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (r *userRepository) GetUserByUsername(username string) (entity.User, error) {
	var user entity.User
	if err := r.db.Where("username = ?", username).Take(&user).Error; err != nil {
		return entity.User{}, err
	}
	return user, nil
}
