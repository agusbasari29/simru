package services

import (
	"github.com/mashingan/smapping"
	"github.com/sip/simru/entity"
	"github.com/sip/simru/repository"
	"github.com/sip/simru/request"
)

type SectionsServices interface {
	GetSections(request request.SectionRequest) (entity.Sections, error)
}

type sectionsService struct {
	sectionRepository repository.SectionRepository
}

func NewSectionsServices(sectionRepository repository.SectionRepository) *sectionsService {
	return &sectionsService{sectionRepository}
}

func (s *sectionsService) GetSections(request request.SectionRequest) (entity.Sections, error) {
	sections := entity.Sections{}
	err := smapping.FillStruct(&sections, smapping.MapFields(&request))
	smapError(err)
	result, err := s.sectionRepository.GetSection(sections)
	if err != nil {
		return result, err
	}
	return result, nil
}