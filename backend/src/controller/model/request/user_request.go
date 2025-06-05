package request

type UserRequest struct {
	Name       string   `json:"name" binding:"required,min=3,max=100"`
	Email      string   `json:"email" binding:"required,email"`
	Password   string   `json:"password" binding:"required,min=6,max=20"`
	Phone      string   `json:"phone" binding:"required,min=11,max=11"`
	Department string   `json:"department" binding:"required"`
	Projects   []string `json:"projects" binding:"required"`
	Enterprise string   `json:"enterprise" binding:"required"`
	Role       string   `json:"role" binding:"required"`
}

type UserUpdateRequest struct {
	Name       string `json:"name" binding:"omitempty,min=3,max=100"`
	Department string `json:"department" binding:"omitempty"`
	Phone      string `json:"phone" binding:"omitempty"`
}
