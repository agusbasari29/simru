package repository

import (
	"github.com/sip/simru/entity"
	"gorm.io/gorm"
)

type sectionRepository struct {
	db *gorm.DB
}

func NewSectionRepository(db *gorm.DB) *sectionRepository {
	return &sectionRepository{db}
}

func (r *sectionRepository) CreateSection(section entity.Sections) (entity.Sections, error) {
	err := r.db.Raw("INSERT INTO sections (section_name, authority, created_at) VALUES (@SectionName, @Authhority, @CreatedAt)", section).Create(&section).Error
	if err != nil {
		return section, err
	}
	return section, nil
}
