package dataaccess

import (
	"log"
	"taskmng/dto"

	"gopkg.in/mgo.v2/bson"
)

//Login - login using email and password
func Login(email string, password string) (dto.User, error) {
	var user dto.User
	err := usersCollection.Find(bson.M{"email": email, "password": password}).One(&user)
	if err != nil {
		log.Print("Error: ", err)
	}
	return user, err
}
