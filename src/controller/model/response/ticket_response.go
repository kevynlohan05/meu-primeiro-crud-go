package response

import "time"

type TicketResponse struct {
	ID            string    `json:"id"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	RequestType   string    `json:"request_type"`
	Priority      string    `json:"priority"`
	AttachmentURL string    `json:"attachment_url"`
	Status        string    `json:"status"`        // Ex: "Pendente", "Em andamento", "Finalizado"
	AsanaTaskID   string    `json:"asana_task_id"` // Útil para vincular com a task criada
	Deadline      time.Time `json:"deadline"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
