package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type TicketEntity struct {
	ID            primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title         string             `bson:"title,omitempty"`
	Description   string             `bson:"description ,omitempty"`
	RequestType   string             `bson:"request_type ,omitempty"`
	Priority      string             `bson:"priority ,omitempty"`
	AttachmentURL string             `bson:"attachment_url ,omitempty"`
}
