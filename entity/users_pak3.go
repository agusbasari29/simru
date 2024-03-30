package entity

import "gorm.io/gorm"

type UserRole string

const (
	Pimpinan UserRole = "pimpinan"
	Admin    UserRole = "admin"
	PIC      UserRole = "pic"
)

type UsersPAK3 struct {
	gorm.Model
	ID       uint `gorm:"primaryKey;autoIncrement"`
	Name     string
	Password string
	UserRole UserRole
}
