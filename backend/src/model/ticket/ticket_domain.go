package model

type ticketDomain struct {
	iD             string
	title          string
	requestUser    string
	sector         string
	description    string
	requestType    string
	priority       string
	attachmentURLs []string
	asanaTaskID    string
	status         string
	comments       []CommentDomain
	projectId      int64
	projectName    string
	asanaProjectID string
}

type CommentDomain struct {
	ID        int64
	TicketID  int64
	Author    string
	Content   string
	CreatedAt int64
}

func (td *ticketDomain) SetProjectName(projectName string) {
	td.projectName = projectName
}
func (td *ticketDomain) GetProjectName() string {
	return td.projectName
}

func (td *ticketDomain) SetAsanaProjectID(asanaProjectID string) {
	td.asanaProjectID = asanaProjectID
}
func (td *ticketDomain) GetAsanaProjectID() string {
	return td.asanaProjectID
}

func (td *ticketDomain) SetID(id string) {
	td.iD = id
}

func (td *ticketDomain) SetStatus(status string) {
	td.status = status
}

func (td *ticketDomain) GetStatus() string {
	return td.status
}

func (td *ticketDomain) GetProjectID() int64 {
	return td.projectId
}

func (td *ticketDomain) SetAsanaTaskID(asanaTaskID string) {
	td.asanaTaskID = asanaTaskID
}

func (td *ticketDomain) GetAsanaTaskID() string {
	return td.asanaTaskID
}

func (td *ticketDomain) GetRequestUser() string {
	return td.requestUser
}

func (td *ticketDomain) GetSector() string {
	return td.sector
}

func (td *ticketDomain) GetID() string {
	return td.iD
}

func (td *ticketDomain) GetTitle() string {
	return td.title
}

func (td *ticketDomain) GetDescription() string {
	return td.description
}

func (td *ticketDomain) GetRequestType() string {
	return td.requestType
}

func (td *ticketDomain) GetPriority() string {
	return td.priority
}

func (td *ticketDomain) GetAttachmentURLs() []string {
	return td.attachmentURLs
}

func (td *ticketDomain) GetComments() []CommentDomain {
	return td.comments
}

func (td *ticketDomain) SetComments(comments []CommentDomain) {
	td.comments = comments
}

func (td *ticketDomain) AddComment(comment CommentDomain) {
	td.comments = append(td.comments, comment)
}

func (td *ticketDomain) SetProjectID(projectID int64) {
	td.projectId = projectID
}
