package main

import (
	"taskmng/utils"
	"taskmng/dataaccess"

	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/middleware/logger"

	"github.com/dgrijalva/jwt-go"
	jwtmiddleware "github.com/iris-contrib/middleware/jwt"
)

func main() {
	app := iris.New()
	app.Use(logger.New())
	dataaccess.InitDbContext()
	utils.InitMailSettings()
	defer dataaccess.TerminateDbContext()

	jwtMiddleware = jwtmiddleware.New(jwtmiddleware.Config{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(AppSecret), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})

	appCors := configCors()

	var authAPI = app.Party("/auth", appCors).AllowMethods(iris.MethodOptions)
	authAPI.Post("/login", login)
	authAPI.Post("/signup", signup)
	authAPI.Post("/verify", verifyEmail)

	// register all tasks API - using jwt token for authentication
	var taskAPI = app.Party("/task", appCors).AllowMethods(iris.MethodOptions)
	taskAPI.Use(jwtMiddleware.Serve)
	taskAPI.Post("/search/{pageSize:int}/{pageIndex:int}", authenticationHandler, searchTask)
	taskAPI.Post("", authenticationHandler, addTask)
	taskAPI.Delete("/{taskId:string}", authenticationHandler, deleteTask)
	taskAPI.Put("/{taskId:string}", authenticationHandler, updateTask)
	taskAPI.Get("/{taskId:string}", authenticationHandler, getTask)

	app.Run(iris.Addr(":8080"))
}

func configCors() context.Handler {
	return cors.New(cors.Options{
		AllowCredentials: true,
		AllowedOrigins:   []string{"*"},
		AllowedHeaders:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	})
}
