package main

import (
	"log"
	"taskmng/dataaccess"
	"taskmng/dto"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	jwtmiddleware "github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris"
)

var jwtMiddleware *jwtmiddleware.Middleware

func authenticationHandler(ctx iris.Context) {
	userToken := jwtMiddleware.Get(ctx)
	if claims, ok := userToken.Claims.(jwt.MapClaims); ok && userToken.Valid {
		var appID = claims[ClaimAppID].(string)
		var email = claims[ClaimEmail].(string)
		var expiredTime, err = time.Parse(ClaimTimeFormat, claims[ClaimExpiredTime].(string))
		log.Print("Current Time: ", time.Now().Format(time.RFC3339))
		log.Print("Token Expired Time: ", expiredTime.Format(time.RFC3339))
		if appID == ApplicationID && err == nil && expiredTime.After(time.Now()) {
			log.Println("User: ", email, " is valid for executing action")
			ctx.Next()
		}
	}
}

func login(ctx iris.Context) {
	var loginInf dto.LoginInfo
	ctx.ReadJSON(&loginInf)
	log.Print("User is trying to login with email: ", loginInf.Email)
	var result dto.LoginResult
	var user, err = dataaccess.Login(loginInf.Email, loginInf.Password)
	if err == nil {
		if user.Status == UserStatusActive {
			log.Print("Login credentials are valid. Going to generate token")
			log.Print("Current Time: ", time.Now().Format(time.RFC3339))
			var expiredTime = time.Now().Add(TokenValidPeriodInMinutes * time.Minute)
			log.Print("Expired Time: ", expiredTime.Format(time.RFC3339))
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				ClaimEmail:       user.Email,
				ClaimAppID:       ApplicationID,
				ClaimExpiredTime: expiredTime,
			})

			// Sign and get the complete encoded token as a string using the secret
			tokenString, _ := token.SignedString([]byte(AppSecret))
			result = dto.LoginResult{IsSuccess: true, Message: ("Welcome: " + user.Email), Token: tokenString}
		} else {
			result = dto.LoginResult{IsSuccess: false, Message: "User status is not active"}
		}
	} else {
		result = dto.LoginResult{IsSuccess: false, Message: "Email or password are invalid"}
	}
	log.Print("Login Result: ", result.Message)

	ctx.JSON(result)
}