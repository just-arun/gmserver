package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type PostModel struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title       string             `json:"title,omitempty" bson:"title,omitempty"`
	Link        string             `json:"link,omitempty",bson:"link,omitempty"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	Keyword     string             `json:"keyword,omitempty" bson:"keyword,omitempty"`
	Body        string             `json:"body,omitempty" bson:"body,omitempty"`
	Author      primitive.ObjectID `json:"author,omitempty" bson:"author,omitempty"`
}
