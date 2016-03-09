package secret

type Credential struct {
	Login    string
	Password string
}

var (
	Blog     = Credential{"blog", "secret-blog-password"}
	CurrTime = Credential{"currtime", "secret-currtime-password"}
	Admin    = Credential{"admin", "secret-admin"}
	User     = Credential{"user", "secret-user"}
)

var credentials = map[Credential]string{
	Blog:     "service",
	CurrTime: "service",
	Admin:    "admin",
	User:     "user",
}

// AuthRole return role for given credentials. Empty role means credentials
// are not valid and user cannot be identified with any role.
func AuthRole(login, password string) string {
	return credentials[Credential{login, password}]
}

// key used to sign JWT tokens
const (
	SigKey = `
		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
	`
)
