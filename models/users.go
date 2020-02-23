package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserModel struct {
	ID       primitive.ObjectID   `json:"id,omitempty" bson:"_id,omitempty"`
	Name     string               `json:"name,omitempty" bson:"name,omitempty"`
	Email    string               `json:"email,omitempty" bson:"email,omitempty"`
	Password string               `json:"password,omitempty" bson:"password,omitempty"`
	PostsID  []primitive.ObjectID `json:"postsId,omitempty" bson:"postsId,omitempty"`
	Posts    []PostModel          `json:"posts,omitempty" bson:"posts,omitempty"`
	Postss   [][]PostModel        `json:"postss,omitempty" bson:"postss,omitempty"`
}
