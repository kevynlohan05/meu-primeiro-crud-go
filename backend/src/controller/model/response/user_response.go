package response

type UserResponse struct {
	ID         string   `json:"id"`
	Name       string   `json:"name"`
	Email      string   `json:"email"`
	Phone      string   `json:"phone"`
	Enterprise string   `json:"enterprise"`
	Department string   `json:"department"`
	Projects   []string `json:"projects"`
	Role       string   `json:"role"` // "admin" ou "user"
}
