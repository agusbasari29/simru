package services

import (
	"github.com/mashingan/smapping"
	"github.com/sip/simru/entity"
	"github.com/sip/simru/repository"
	"github.com/sip/simru/request"
)

type UserRolesService interface {
	GetUserRoles(request request.UserRolesRequest) (entity.UserRoles, error)
}

type userRolesService struct {
	userRolesRepository repository.UserRoleRepository
}

func NewUserRolesServices(userRolesRepository repository.UserRoleRepository) *userRolesService {
	return &userRolesService{userRolesRepository}
}

func (s *userRolesService) GetUserRoles(request request.UserRolesRequest) (entity.UserRoles, error) {
	userRoles := entity.UserRoles{}
	err := smapping.FillStruct(&userRoles, smapping.MapFields(&request))
	smapError(err)
	result, err := s.userRolesRepository.GetUserRoles(userRoles)
	if err != nil {
		return result, err
	}
	return result, nil
}
