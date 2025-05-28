package entity


type TicketEntity struct {
	ID             int64    `json:"id" gorm:"primaryKey;autoIncrement"`
	Title          string   `json:"title"`
	RequestUser    string   `json:"request_user"`
	Sector         string   `json:"sector"`
	Description    string   `json:"description"`
	RequestType    string   `json:"request_type"`
	Priority       string   `json:"priority"`
	AttachmentURLs string   `json:"attachment_urls"` // armazenaremos como JSON
	Department     string   `bson:"department"`
	Status         string   `json:"status"`
	AsanaTaskID    string   `json:"asana_task_id"`
	Projects       string   `json:"projects"` // pode ser JSON ou string simples
	Comments       string   `json:"comments"` // salva como JSON (serializado)
}

type Comment struct {
	Author    string `bson:"author" json:"author"`
	Message   string `bson:"message" json:"message"`
	Timestamp int64  `bson:"timestamp" json:"timestamp"`
}
