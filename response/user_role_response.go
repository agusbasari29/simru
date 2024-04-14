package response

import "github.com/sip/simru/entity"

type UserRolesResponse struct {
	ID       uint64          `json:"id"`
	Role     string          `json:"role"`
	RoleName string          `json:"role_name"`
	Sections SectionsResponse `json:"section"`
	// SectionID uint64 `json:"section_id"`
}

func UserRolesResponseFormatter(userRoles entity.UserRoles, sections entity.Sections) UserRolesResponse {
	formatter := UserRolesResponse{}
	formatter.ID = userRoles.ID
	formatter.Role = userRoles.Role
	formatter.RoleName = userRoles.RoleName
	formatter.Sections = SectionsResponseFormatter(sections)
	return formatter
}