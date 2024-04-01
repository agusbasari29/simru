package entity

import "gorm.io/gorm"

type Sections struct {
	gorm.Model
	ID          uint64 `gorm:"primaryKey;autoIncrement"`
	SectionName string
}
