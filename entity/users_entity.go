package entity

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	ID         uint64 `gorm:"primaryKey;autoIncrement"`
	Username   string
	Password   string
	PersonID   uint64
	Person     Persons `gorm:"foreignKey:PersonID"`
	UserRoleID uint64
	UserRole   UserRoles `gorm:"foreignKey:UserRoleID"`
}
