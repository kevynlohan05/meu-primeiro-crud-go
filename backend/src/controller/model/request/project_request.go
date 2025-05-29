package request

type ProjectRequest struct {
	Name    string `json:"name" binding:"required,min=1,max=100"`
	IdAsana string `json:"id_asana" binding:"required,min=1,max=100"`
}

type ProjectUpdateRequest struct {
	Name    string `json:"name" binding:"omitempty,min=1,max=100"`
	IdAsana string `json:"id_asana" binding:"omitempty,min=1,max=100"`
}
