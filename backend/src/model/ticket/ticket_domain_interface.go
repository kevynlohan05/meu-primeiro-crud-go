package model

type TicketDomainInterface interface {
	GetID() string
	GetTitle() string
	GetRequestUser() string
	GetSector() string
	GetDescription() string
	GetRequestType() string
	GetPriority() string
	GetAttachmentURL() string
	GetAsanaTaskID() string
	GetStatus() string
	GetComments() []CommentDomain
	GetProjects() string

	SetComments([]CommentDomain)
	SetStatus(string)
	SetAsanaTaskID(string)
	SetID(string)

	AddComment(CommentDomain)
}

func NewTicketDomain(title, requestUser, sector, description, requestType, priority, attachmentURL, projects string) TicketDomainInterface {
	return &ticketDomain{
		title:         title,
		requestUser:   requestUser,
		sector:        sector,
		description:   description,
		requestType:   requestType,
		priority:      priority,
		attachmentURL: attachmentURL,
		projects:      projects,
	}
}

func NewTicketUpdateDomain(title, description, requestType, priority, attachmentURL, status string) TicketDomainInterface {
	return &ticketDomain{
		title:         title,
		description:   description,
		requestType:   requestType,
		priority:      priority,
		attachmentURL: attachmentURL,
		status:        status,
	}
}
