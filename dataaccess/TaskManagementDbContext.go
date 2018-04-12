package dataaccess

import (
	"gopkg.in/mgo.v2"
)

var session *mgo.Session
var tasksCollection *mgo.Collection
var usersCollection *mgo.Collection

//InitDbContext used for initializing database context
func InitDbContext() {
	// Database connection
	session, err := mgo.Dial("localhost:27017")
	if nil != err {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)

	tasksCollection = session.DB("task_management").C("Task")
	usersCollection = session.DB("task_management").C("User")
}

//TerminateDbContext used for terminating database context
func TerminateDbContext() {
	session.Close()
}
