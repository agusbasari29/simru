package response

import "github.com/sip/simru/entity"

type ResponseCredential struct {
	Token     string           `json:"token"`
	UserID    uint64           `json:"user_id"`
	Role      entity.UserRoles `json:"role"`
	Person    entity.Persons   `json:"person"`
	Issuer    string           `json:"issuer"`
	IssuedAt  int64            `json:"issued_at"`
	ExpiresAt int64            `json:"expired_at"`
}
