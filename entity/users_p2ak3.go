package entity

import "gorm.io/gorm"

type UsersP2AK3 struct {
	gorm.Model
	ID       uint `gorm:"primaryKey;autoIncrement"`
	Name     string
	Password string
	UserRole UserRole
}
