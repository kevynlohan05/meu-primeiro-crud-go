package model

type TicketDomainInterface interface {
	GetTitle() string
	GetDescription() string
	GetRequestType() string
	GetPriority() string
	GetAttachmentURL() string
	GetID() string
	GetUserEmail() string

	SetID(string)
}

func NewTicketDomain(title, description, requestType, priority, attachmentURL, userEmail string) TicketDomainInterface {
	return &ticketDomain{
		title:         title,
		description:   description,
		requestType:   requestType,
		priority:      priority,
		attachmentURL: attachmentURL,
		userEmail:     userEmail,
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
