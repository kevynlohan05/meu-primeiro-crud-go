package request

type UserRequest struct {
	Name       string `json:"name" binding:"required,min=3,max=100"`
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required,min=6,max=20"`
	Department string `json:"department" binding:"required"`
	Role       string `json:"role" binding:"required"` // "admin" ou "user"
}

type UserUpdateRequest struct {
	Name       string `json:"name" binding:"omitempty,min=3,max=100"`
	Department string `json:"department" binding:"omitempty"`
}
