package entity

import (
	ticketModel "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/ticket"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TicketEntity struct {
	ID            primitive.ObjectID          `json:"id" bson:"_id,omitempty"`
	Title         string                      `bson:"title,omitempty"`
	RequestUser   string                      `bson:"request_user,omitempty"`
	Sector        string                      `bson:"sector,omitempty"`
	Description   string                      `bson:"description,omitempty"`
	RequestType   string                      `bson:"request_type,omitempty"`
	Priority      string                      `bson:"priority,omitempty"`
	AttachmentURL string                      `bson:"attachment_url,omitempty"`
	Status        string                      `bson:"status,omitempty"`
	AsanaTaskID   string                      `bson:"asana_task_id,omitempty"`
	Comments      []ticketModel.CommentDomain `bson:"comments,omitempty"`
}

type Comment struct {
	Author    string `bson:"author" json:"author"`
	Message   string `bson:"message" json:"message"`
	Timestamp int64  `bson:"timestamp" json:"timestamp"`
}
