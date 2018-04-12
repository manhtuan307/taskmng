package dto

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// A User represents a user in system
type User struct {
	ID             bson.ObjectId `bson:"_id,omitempty"`
	Email          string        `bson:"email"`
	Password       string        `bson:"password"`
	CreatedTime    time.Time     `bson:"createdTime"`
	Status         int           `bson:"status"`
	ActivationCode string        `bson:"activationCode"`
}
