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
	Projects       string            `json:"projects"`
}

type CommentResponse struct {
	Author    string `json:"author"`
	Message   string `json:"message"`
	Timestamp int64  `json:"timestamp"`
}
