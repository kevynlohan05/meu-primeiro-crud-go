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
	projects       string
}

type CommentDomain struct {
	Author    string
	Message   string
	Timestamp int64
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

func (td *ticketDomain) GetProjects() string {
	return td.projects
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
