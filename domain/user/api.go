package user

type User struct {
	UserID        string `json:"user_id"`
	Email         string `json:"email"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Nickname      string `json:"nickname"`
	LastIP        string `json:"last_ip"`
	LoginsCount   int    `json:"logins_count"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
	LastLogin     string `json:"last_login"`
	EmailVerified bool   `json:"email_verified"`
}
