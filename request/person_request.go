package request

type PersonRequest struct {
	ID             uint64 `json:"id"`
	Name           string `json:"name"`
	NIP            string `json:"nip"`
	Email          string `json:"email"`
	Phone          string `json:"phone"`
	CompanyName    string `json:"company_name"`
	CompanyAddress string `json:"company_address"`
}
