package response

type PersonResponse struct {
	ID             uint64 `json:"id"`
	Name           string `json:"name"`
	NIP            string `json:"nip"`
	CompanyName    string `json:"company_name"`
	CompanyAddress string `json:"company_address"`
	Email          string `json:"email"`
	Phone          string `json:"phone"`
}
