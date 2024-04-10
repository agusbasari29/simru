package repository

import (
	"github.com/sip/simru/entity"
	"gorm.io/gorm"
)

type SectionRepository interface {
	CreateSection(section entity.Sections) (entity.Sections, error)
	UpdateSection(section entity.Sections) (entity.Sections, error)
	GetSection(section entity.Sections) (entity.Sections, error)
	GetSections() ([]entity.Sections, error)
	DeleteSection(section entity.Sections) bool
}

type sectionRepository struct {
	db *gorm.DB
}

func NewSectionRepository(db *gorm.DB) *sectionRepository {
	return &sectionRepository{db}
}

func (r *sectionRepository) CreateSection(section entity.Sections) (entity.Sections, error) {
	err := r.db.Create(&section).Error
	if err != nil {
		return section, err
	}
	return section, nil
}

func (r *sectionRepository) UpdateSection(section entity.Sections) (entity.Sections, error) {
	err := r.db.Model(&entity.Sections{ID: section.ID}).Save(&section).Error
	if err != nil {
		return section, err
	}
	return section, nil
}

func (r *sectionRepository) GetSection(section entity.Sections) (entity.Sections, error) {
	err := r.db.First(&section).Error
	if err != nil {
		return section, err
	}
	return section, nil
}

func (r *sectionRepository) GetSections() ([]entity.Sections, error) {
	section := []entity.Sections{}
	err := r.db.Find(&section).Error
	if err != nil {
		return section, err
	}
	return section, nil
}

func (r *sectionRepository) DeleteSection(section entity.Sections) bool {
	return r.db.Delete(&section).Error == nil
}
