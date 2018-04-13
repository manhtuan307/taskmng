package dto

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// A Task represent task and errand
type Task struct {
	ID      bson.ObjectId `bson:"_id,omitempty"`
	UserID  bson.ObjectId `bson:"userID,omitempty"`
	Name    string        `bson:"name"`
	Status  string        `bson:"status"`
	Created time.Time     `bson:"created"`
	Updated time.Time     `bson:"updated"`
}
