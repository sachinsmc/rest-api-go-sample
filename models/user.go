package models

import "gopkg.in/mgo.v2/bson"

/*
User Model
Represents a User, we uses bson keyword to tell the mgo driver how to name
the properties in mongodb document
*/
type User struct {
	ID        bson.ObjectId `bson:"_id" json:"id"`
	FirstName string        `bson:"first_name" json:"first_name"`
	LastName  string        `bson:"last_name" json:"last_name"`
}
