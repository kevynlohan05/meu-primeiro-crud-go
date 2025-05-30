package model

type TicketDomainInterface interface {
	GetID() string
	GetTitle() string
	GetRequestUser() string
	GetSector() string
	GetDescription() string
	GetRequestType() string
	GetPriority() string
	GetAttachmentURLs() []string
	GetAsanaTaskID() string
	GetStatus() string
	GetComments() []CommentDomain
	GetProjectID() int64
	GetAsanaProjectID() string
	GetProjectName() string

	SetProjectName(string)
	SetAsanaProjectID(string)
	SetProjectID(int64)
	SetComments([]CommentDomain)
	SetStatus(string)
	SetAsanaTaskID(string)
	SetID(string)

	AddComment(CommentDomain)
}

func NewTicketDomain(title, requestUser, sector, description, requestType, priority, projectName string, attachmentURL []string) TicketDomainInterface {
	return &ticketDomain{
		title:          title,
		requestUser:    requestUser,
		sector:         sector,
		description:    description,
		requestType:    requestType,
		priority:       priority,
		attachmentURLs: attachmentURL,
		projectName:    projectName,
	}
}

func NewTicketUpdateDomain(title, description, requestType, priority, status string) TicketDomainInterface {
	return &ticketDomain{
		title:       title,
		description: description,
		requestType: requestType,
		priority:    priority,
		status:      status,
	}
}

func NewTicketDomainFromEntity(title, requestUser, sector, description, requestType, priority string, attachmentURL []string, projectId int64) TicketDomainInterface {
	return &ticketDomain{
		title:          title,
		requestUser:    requestUser,
		sector:         sector,
		description:    description,
		requestType:    requestType,
		priority:       priority,
		attachmentURLs: attachmentURL,
		projectId:      projectId,
	}
}
