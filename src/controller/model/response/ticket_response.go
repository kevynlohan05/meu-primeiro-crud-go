package response

type TicketResponse struct {
	ID            string `json:"id"`
	Title         string `json:"title"`
	RequestUser   string `json:"request_user"`
	Sector        string `json:"sector"`
	Description   string `json:"description"`
	RequestType   string `json:"request_type"`
	Priority      string `json:"priority"`
	AttachmentURL string `json:"attachment_url"`
	AsanaTaskID   string `json:"asana_task_id"`
	Status        string `json:"status"`
}
