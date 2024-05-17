package entity

import "gorm.io/gorm"

type Status string

const (
	None    Status = "none"
	BalaiK3 Status = "balai_k3"
	PJK3    Status = "pjk3"
)

type Persons struct {
	gorm.Model
	ID             uint64 `gorm:"primaryKey;autoIncrement"`
	Name           string
	NIP            string `gorm:"column:nip"`
	CompanyName    string
	CompanyAddress string
	Status         Status
	Email          string
	Phone          string

	//...
}
