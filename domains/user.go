package domains

// User domain
type User struct {
	ID          string   `json:"id"`
	Username    string   `json:"username"`
	Password    string   `json:"-"`
	Email       string   `json:"email"`
	Identifiers []string `json:"identifiers"`
}
