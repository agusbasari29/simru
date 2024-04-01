package entity

import "gorm.io/gorm"

type Persons struct {
	gorm.Model
	ID          uint64 `gorm:"primaryKey;autoIncrement"`
	Name        string
	NIP         string
	CompanyName string
	//...
}
