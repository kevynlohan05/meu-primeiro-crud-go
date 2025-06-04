package response

type TicketResponse struct {
	ID             string            `json:"id"`
	Status         string            `json:"status"`
	Title          string            `json:"title"`
	RequestUser    string            `json:"request_user"`
	Sector         string            `json:"sector"`
	Description    string            `json:"description"`
	RequestType    string            `json:"request_type"`
	Priority       string            `json:"priority"`
	AttachmentURLs []string          `json:"attachment_url"`
	AsanaTaskID    string            `json:"asana_task_id"`
	Comments       []CommentResponse `json:"comments"`
	ProjectName    string            `json:"project_name"`
}

type CommentResponse struct {
	ID        int64  `json:"id"`
	TicketID  int64  `json:"ticket_id"`
	Author    string `json:"author"`
	Content   string `json:"content"`
	CreatedAt int64  `json:"created_at"`
}
