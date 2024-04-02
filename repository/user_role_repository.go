package repository

import (
	"github.com/sip/simru/entity"
	"gorm.io/gorm"
)

type userRoleRepository struct {
	db *gorm.DB
}

func NewUserRoleRepository(db *gorm.DB) *userRoleRepository {
	return &userRoleRepository{db}
}

func (r *userRoleRepository) CreateRole(role entity.UserRoles) (entity.UserRoles, error) {
	err := r.db.Raw("INSERT INTO user_roles (roles, role_name, section_id, created_at) VALUE (@Role, @RoleName, @SectionID, @CreatedAt)", role).Create(&role).Error
	if err != nil {
		return role, err
	}
	return role, nil
}
