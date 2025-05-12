package model

type ticketDomain struct {
	iD            string
	title         string
	description   string
	requestType   string
	priority      string
	attachmentURL string
	userEmail     string
}

func (td *ticketDomain) SetID(id string) {
	td.iD = id
}

func (td *ticketDomain) GetUserEmail() string {
	return td.userEmail
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

func (td *ticketDomain) GetAttachmentURL() string {
	return td.attachmentURL
}
