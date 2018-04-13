package main

import (
	"log"
	"taskmng/dataaccess"
	"taskmng/dto"
	"taskmng/utils"
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
		if user.Status == dto.UserStatusActive {
			log.Print("Login credentials are valid. Going to generate token")
			log.Print("Current Time: ", time.Now().Format(time.RFC3339))
			var expiredTime = time.Now().Add(TokenValidPeriodInMinutes * time.Minute)
			log.Print("Expired Time: ", expiredTime.Format(time.RFC3339))
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				ClaimUserID:      user.ID.Hex(),
				ClaimEmail:       user.Email,
				ClaimAppID:       ApplicationID,
				ClaimExpiredTime: expiredTime,
			})

			// Sign and get the complete encoded token as a string using the secret
			tokenString, _ := token.SignedString([]byte(AppSecret))
			result = dto.LoginResult{IsSuccess: true, Message: ("Welcome: " + user.Email),
				Token: tokenString, ExpiredTime: expiredTime.Format(time.RFC3339),
				UserID: user.ID.Hex()}
		} else {
			result = dto.LoginResult{IsSuccess: false, Message: "User status is not active"}
		}
	} else {
		result = dto.LoginResult{IsSuccess: false, Message: "Email or password are invalid"}
	}
	log.Print("Login Result: ", result.Message)

	ctx.JSON(result)
}

func signup(ctx iris.Context) {
	var registraionInfo dto.RegistrationInfo
	ctx.ReadJSON(&registraionInfo)
	var result dto.ActionResult
	if registraionInfo.Email != "" && registraionInfo.Password != "" {
		if registraionInfo.Password == registraionInfo.ConfirmPassword {
			var user, err = dataaccess.Register(registraionInfo.Email, registraionInfo.Password)
			if err == nil {
				sendConfirmMail(user)
				result = dto.ActionResult{IsSuccess: true, Message: "Thank you for your registration. Please check verification email to verify your email address"}
			} else {
				result = dto.ActionResult{IsSuccess: false, Message: err.Error()}
			}
		} else {
			result = dto.ActionResult{IsSuccess: false, Message: "Password and Confirm Password mismatched"}
		}
	} else {
		result = dto.ActionResult{IsSuccess: false, Message: "Registration information is invalid"}
	}
	ctx.JSON(result)
}

func verifyEmail(ctx iris.Context) {
	var verifyInfo dto.EmailVerificationInfo
	ctx.ReadJSON(&verifyInfo)
	var result dto.ActionResult
	err := dataaccess.VerifyRegistration(verifyInfo.Email, verifyInfo.VerifyCode)
	if err == nil {
		result = dto.ActionResult{IsSuccess: true, Message: "Email has been verified"}
	} else {
		result = dto.ActionResult{IsSuccess: false, Message: err.Error()}
	}
	ctx.JSON(result)
}

func sendConfirmMail(user dto.User) {
	var verifyLink = "http://174.16.10.107/taskmng/verifyRegistration.html?email=" + user.Email + "&code=" + user.ActivationCode
	var content = "Dear Sir/Madam. Thank you for your registration." +
		"Please click the following link bellow to verify your email for Task Management registration: " +
		verifyLink
	utils.SendMailToOne(user.Email, content)
}
