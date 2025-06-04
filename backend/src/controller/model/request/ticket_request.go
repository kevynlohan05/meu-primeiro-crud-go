package request

type TicketRequest struct {
	Title       string `json:"title" binding:"required,min=3,max=100"`
	Description string `json:"description" binding:"required,min=3,max=1000"`
	RequestType string `json:"request_type" binding:"required"` // Ex: "Bugs", "Dúvidas", etc.
	Priority    string `json:"priority" binding:"required"`     // Ex: "Baixa", "Média", "Alta"
	Project     string `json:"project" binding:"required"`
}

type TicketUpdateRequest struct {
	Title         string `json:"title" binding:"omitempty,min=3,max=100"`
	RequestUser   string `json:"request_user" binding:"omitempty,min=3,max=100"`
	Sector        string `json:"sector" binding:"omitempty"`
	Description   string `json:"description" binding:"omitempty,min=3,max=1000"`
	RequestType   string `json:"request_type" binding:"omitempty"` // Ex: "Suporte", "Financeiro", etc.
	Priority      string `json:"priority" binding:"omitempty"`     // Ex: "Baixa", "Média", "Alta"
	AttachmentURL string `json:"attachment_url,omitempty"`         // Ex: link para o arquivo no S3, etc.
	Status        string `json:"status" binding:"omitempty"`
}

type AddCommentRequest struct {
	Content string `json:"content" binding:"required"`
}
