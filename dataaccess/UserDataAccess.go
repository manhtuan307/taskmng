package dataaccess

import (
	"log"
	"taskmng/dto"
	"taskmng/utils"
	"time"

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

//Register - register using email and password
func Register(email string, password string) (dto.User, error) {
	var user = dto.User{Email: email, Password: password, Status: 0,
		CreatedTime: time.Now(), ActivationCode: utils.RandStringBytes(8)}
	err := usersCollection.Insert(user)
	if err != nil {
		log.Print("Error: ", err)
	}
	return user, err
}

//VerifyRegistration - verify user registration
func VerifyRegistration(email string, verifyCode string) error {
	var user dto.User
	err := usersCollection.Find(bson.M{"email": email}).One(&user)
	if err == nil {
		if user.ActivationCode == verifyCode {
			if user.Status == dto.UserStatusInactive {
				var query = bson.M{"_id": user.ID.Hex}
				var change = bson.M{"$set": bson.M{"status": dto.UserStatusActive}}
				err := usersCollection.Update(query, change)
				if err != nil {
					log.Print("Error: ", err)
				}
			} else {
				panic("Registration has already confirm")
			}
		} else {
			panic("Verify code is wrong")
		}
	}
	return err
}
