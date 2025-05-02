package request

type UserRequest struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Department string `json:"department"`
	Role       string `json:"role"` // "admin" ou "user"
}
