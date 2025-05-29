package entity

type ProjectEntity struct {
	ID      int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name    string `json:"name"`
	IdAsana string `json:"id_asana"`
}
