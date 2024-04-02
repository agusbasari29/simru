package repository

import (
	"github.com/sip/simru/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) CreateUser(user entity.Users) (entity.Users, error) {
	err := r.db.Raw("INSERT INTO users (username, password, user_role_id, created_at) VALUE (@Username, @Password, @UserRoleID, @CreatedAt)", user).Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}