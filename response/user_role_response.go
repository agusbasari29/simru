package response

type UserRoleResponse struct {
	ID       uint64          `json:"id"`
	Role     string          `json:"role"`
	RoleName string          `json:"role_name"`
	Sections SectionResponse `json:"section"`
	// SectionID uint64 `json:"section_id"`
}
