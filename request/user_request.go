package request

type UserRequest struct {
	ID uint64 `json:"id"`
}

type UserRegisterRequest struct {
	Username       string `json:"username" validate:"required,alphanum"`
	Password       string `json:"password" validate:"required"`
	UserRoleID     uint64 `json:"user_role_id" validate:"required"`
	Name           string `json:"name" validate:"required"`
	NIP            string `json:"nip" validate:"required"`
	Email          string `json:"email" validate:"required"`
	Phone          string `json:"phone" validate:"required"`
	CompanyName    string `json:"company_name" validate:"required"`
	CompanyAddress string `json:"company_address" validate:"required"`
	// Person   PersonRequest
}
