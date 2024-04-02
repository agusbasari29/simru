package request

type UserRegisterRequest struct {
	Name           string `json:"name"`
	NIP            string `json:"nip"`
	Username       string `json:"username"`
	Password       string `json:"password"`
	Email          string `json:"email"`
	Phone          string `json:"phone"`
	CompanyName    string `json:"company_name"`
	CompanyAddress string `json:"company_address"`
	RoleID         uint64 `json:"role_id"`
	SectionID      uint64 `json:"section_id"`
}

