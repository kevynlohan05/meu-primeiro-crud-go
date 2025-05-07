package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type TicketEntity struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title         string             `bson:"title" json:"title"`
	Description   string             `bson:"description" json:"description"`
	RequestType   string             `bson:"request_type" json:"request_type"`
	Priority      string             `bson:"priority" json:"priority"`
	AttachmentURL string             `bson:"attachment_url" json:"attachment_url"`
}
