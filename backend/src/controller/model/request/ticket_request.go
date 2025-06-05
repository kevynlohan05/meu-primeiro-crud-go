package request

type TicketRequest struct {
	Title       string `json:"title" binding:"required,min=3,max=100"`
	Description string `json:"description" binding:"required,min=3,max=1000"`
	RequestType string `json:"request_type" binding:"required"`
	Priority    string `json:"priority" binding:"required"`
	Project     string `json:"project" binding:"required"`
}

type TicketUpdateRequest struct {
	Title         string `json:"title" binding:"omitempty,min=3,max=100"`
	RequestUser   string `json:"request_user" binding:"omitempty,min=3,max=100"`
	Sector        string `json:"sector" binding:"omitempty"`
	Description   string `json:"description" binding:"omitempty,min=3,max=1000"`
	RequestType   string `json:"request_type" binding:"omitempty"`
	Priority      string `json:"priority" binding:"omitempty"`
	AttachmentURL string `json:"attachment_url,omitempty"`
	Status        string `json:"status" binding:"omitempty"`
}

type AddCommentRequest struct {
	Content string `json:"content" binding:"required"`
}

type UpdateCommentRequest struct {
	Content string `json:"content" binding:"required"`
}
