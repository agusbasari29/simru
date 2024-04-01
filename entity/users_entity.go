package entity

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	ID         uint64 `gorm:"primaryKey;autoIncrement"`
	Username   string
	Password   string
	PersonID   Persons
	UserRoleID UserRoles
}
