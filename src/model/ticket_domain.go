package model

type TicketDomainInterface interface {
	GetTitle() string
	GetDescription() string
	GetRequestType() string
	GetPriority() string
	GetAttachmentURL() string
}

func NewTicketDomain(title, description, requestType, priority, attachmentURL string) TicketDomainInterface {
	return &ticketDomain{title, description, requestType, priority, attachmentURL}
}

type ticketDomain struct {
	title         string
	description   string
	requestType   string
	priority      string
	attachmentURL string
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

func (td *ticketDomain) GetAttachmentURL() string {
	return td.attachmentURL
}
