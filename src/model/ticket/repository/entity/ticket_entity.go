package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type TicketEntity struct {
	ID            primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title         string             `bson:"title,omitempty"`
	RequestUser   string             `bson:"request_user,omitempty"`
	Sector        string             `bson:"sector,omitempty"`
	Description   string             `bson:"description,omitempty"`
	RequestType   string             `bson:"request_type,omitempty"`
	Priority      string             `bson:"priority,omitempty"`
	AttachmentURL string             `bson:"attachment_url,omitempty"`
	Status        string             `bson:"status,omitempty"`
	AsanaTaskID   string             `bson:"asana_task_id,omitempty"`
}
