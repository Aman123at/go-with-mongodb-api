package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name    string             `json:"name"`
	Age     int                `json:"age"`
	Email   string             `json:"email"`
	Address string             `json:"address"`
	Rating  int                `json:"rating"`
}

func (u *User) IsEmpty() bool {
	isEmpty := false
	if u.Name == "" || u.Email == "" {
		isEmpty = true
	}
	return isEmpty
}
