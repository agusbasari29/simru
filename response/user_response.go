package response

import (
	"time"

	"github.com/sip/simru/entity"
	"gorm.io/gorm"
)

type UserResponse struct {
	ID        uint64           `json:"id"`
	Username  string           `json:"username"`
	UserRole  UserRolesResponse `json:"user_role"`
	Person    PersonResponse   `json:"person"`
	CreatedAt time.Time        `json:"created_at"`
	UpdatedAt time.Time        `json:"updated_at"`
	DeletedAt gorm.DeletedAt   `json:"deleted_at"`
}

type UserDataResponse struct {
	User       interface{} `json:"user_data"`
	Credential interface{} `json:"credential"`
}

func UserResponseFormatter(user entity.Users, userRoles entity.UserRoles, person entity.Persons, sections entity.Sections) UserResponse {
	formatter := UserResponse{}
	formatter.ID = user.ID
	formatter.Username = user.Username
	formatter.UserRole = UserRolesResponseFormatter(userRoles, sections)
	formatter.Person = PersonResponseFormatter(person)
	formatter.CreatedAt = user.CreatedAt
	formatter.UpdatedAt = user.UpdatedAt
	formatter.DeletedAt = user.DeletedAt

	return formatter
}

func UserDataResponseFormatter(user interface{}, credential interface{}) UserDataResponse {
	userData := UserDataResponse{
		User:       user,
		Credential: credential,
	}

	return userData
}
