package entity

import "gorm.io/gorm"

type Persons struct {
	gorm.Model
	ID             uint64 `gorm:"primaryKey;autoIncrement"`
	Name           string
	NIP            string `gorm:"column:nip"`
	CompanyName    string
	CompanyAddress string
	Email          string
	Phone          string

	//...
}
