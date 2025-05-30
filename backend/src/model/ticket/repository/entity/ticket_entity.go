package entity

type TicketEntity struct {
	ID             int64  `json:"id" gorm:"primaryKey;autoIncrement"`
	Title          string `json:"title"`
	RequestUser    string `json:"request_user"`
	Sector         string `json:"sector"`
	Description    string `json:"description"`
	RequestType    string `json:"request_type"`
	Priority       string `json:"priority"`
	AttachmentURLs string `json:"attachment_urls"` // armazenado como JSON ou string separada
	Status         string `json:"status"`
	AsanaTaskID    string `json:"asana_task_id"`
	ProjectID      int64  `json:"project_id"`
	Comments       string `json:"comments"`
}

type Comment struct {
	Author    string `json:"author"`
	Message   string `json:"message"`
	Timestamp int64  `json:"timestamp"`
}
