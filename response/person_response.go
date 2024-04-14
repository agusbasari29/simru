package response

import (
	"github.com/sip/simru/entity"
)

type PersonResponse struct {
	ID             uint64 `json:"id"`
	Name           string `json:"name"`
	NIP            string `json:"nip"`
	CompanyName    string `json:"company_name"`
	CompanyAddress string `json:"company_address"`
	Email          string `json:"email"`
	Phone          string `json:"phone"`
}

func PersonResponseFormatter(person entity.Persons) PersonResponse {
	formatter := PersonResponse{}
	formatter.ID = person.ID
	formatter.Name = person.Name
	formatter.NIP = person.NIP
	formatter.CompanyName = person.CompanyName
	formatter.CompanyAddress = person.CompanyAddress
	formatter.Email = person.Email
	formatter.Phone = person.Phone
	return formatter
}
