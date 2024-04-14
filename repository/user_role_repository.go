package repository

import (
	"github.com/sip/simru/entity"
	"gorm.io/gorm"
)

type UserRoleRepository interface {
	CreateUserRole(role entity.UserRoles) (entity.UserRoles, error)
	UpdateUserRole(role entity.UserRoles) (entity.UserRoles, error)
	GetUserRoles(role entity.UserRoles) (entity.UserRoles, error)
	GetAllUserRoles() ([]entity.UserRoles, error)
	DeleteUserRole(role entity.UserRoles) bool
}

type userRoleRepository struct {
	db *gorm.DB
}

func NewUserRoleRepository(db *gorm.DB) *userRoleRepository {
	return &userRoleRepository{db}
}

func (r *userRoleRepository) CreateUserRole(role entity.UserRoles) (entity.UserRoles, error) {
	err := r.db.Raw("INSERT INTO user_roles (role, role_name, section_id, created_at) VALUE (@Role, @RoleName, @SectionID, @CreatedAt)", role).Create(&role).Error
	if err != nil {
		return role, err
	}
	return role, nil
}

func (r *userRoleRepository) UpdateUserRole(role entity.UserRoles) (entity.UserRoles, error) {
	err := r.db.Model(&entity.UserRoles{ID: role.ID}).Save(&role).Error
	if err != nil {
		return role, err
	}
	return role, nil
}

func (r *userRoleRepository) GetUserRoles(userRole entity.UserRoles) (entity.UserRoles, error) {
	err := r.db.First(&userRole).Error
	if err != nil {
		return userRole, err
	}
	return userRole, nil
}

func (r *userRoleRepository) GetAllUserRoles() ([]entity.UserRoles, error) {
	roles := []entity.UserRoles{}
	err := r.db.Find(&roles).Error
	if err != nil {
		return roles, err
	}
	return roles, nil
}

func (r *userRoleRepository) DeleteUserRole(role entity.UserRoles) bool {
	err := r.db.Delete(&role).Error
	return err == nil
}
