package entity


type UserEntity struct {
	ID         int      `json:"id" gorm:"primaryKey;autoIncrement"`
	Name       string   `json:"name"`
	Email      string   `json:"email"`
	Password   string   `json:"password"`
	Phone      string   `json:"phone"`
	Enterprise string   `json:"enterprise"`
	Department string   `json:"department"`
	Projects   string   `json:"projects"` // pode armazenar como JSON
	Role       string   `json:"role"`
}
