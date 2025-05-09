package model

type TicketDomainInterface interface {
	GetTitle() string
	GetDescription() string
	GetRequestType() string
	GetPriority() string
	GetAttachmentURL() string
	GetID() string

	SetID(string)
}

func NewTicketDomain(title, description, requestType, priority, attachmentURL string) TicketDomainInterface {
	return &ticketDomain{
		title:         title,
		description:   description,
		requestType:   requestType,
		priority:      priority,
		attachmentURL: attachmentURL,
	}
}

func NewTicketUpdateDomain(title, description, requestType, priority, attachmentURL string) TicketDomainInterface {
	return &ticketDomain{
		title:         title,
		description:   description,
		requestType:   requestType,
		priority:      priority,
		attachmentURL: attachmentURL,
	}
}
