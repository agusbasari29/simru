package response

type ResponseCredential struct {
	Token  string `json:"token"`
	UserID uint64 `json:"user_id"`
	// Role      UserRolesResponse `json:"role"`
	// Person    PersonResponse    `json:"person"`
	Issuer    string `json:"issuer"`
	IssuedAt  int64  `json:"issued_at"`
	ExpiresAt int64  `json:"expired_at"`
}
