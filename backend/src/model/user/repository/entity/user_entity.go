package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserEntity struct {
	ID         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name       string             `bson:"name,omitempty"`
	Email      string             `bson:"email,omitempty"`
	Password   string             `bson:"password,omitempty"`
	Phone      string             `bson:"phone,omitempty"`
	Enterprise string             `bson:"enterprise,omitempty"`
	Department string             `bson:"department,omitempty"`
	Projects   []string           `bson:"projects,omitempty"`
	Role       string             `bson:"role,omitempty"`
}
