package main

import (
	"log"
	"taskmng/dto"

	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func main() {
	app := iris.Default()
	app.Use(logger.New())

	appCors := cors.New(cors.Options{
		AllowCredentials: true,
		AllowedOrigins:   []string{"*"},
	})
	app.Use(appCors)

	// Database connection
	session, err := mgo.Dial("localhost:27017")
	if nil != err {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	tasksCollection := session.DB("task_management").C("Task")

	app.Get("/task/{pageSize:int}/{pageIndex:int}", func(ctx iris.Context) {
		var tasks []dto.Task

		var pageSize, _ = ctx.Params().GetInt("pageSize")
		if pageSize < 1 {
			pageSize = PAGE_SIZE
		}
		var pageIndex, _ = ctx.Params().GetInt("pageIndex")
		if pageIndex < 1 {
			pageIndex = 1
		}
		var tasksCount, err = tasksCollection.Find(bson.M{}).Count()
		var numOfPages = tasksCount / pageSize
		var numSkip = (pageIndex - 1) * pageSize

		err = tasksCollection.Find(bson.M{}).Limit(pageSize).Skip(numSkip).All(&tasks)
		if err != nil {
			log.Fatal(err)
		}
		var result = dto.ActionResult{true, ""}
		var getTaskResult = dto.GetTasksActionResult{Result: &result, Tasks: tasks, NumOfPages: numOfPages, PageIndex = pageIndex}
		ctx.JSON(getTaskResult)
	})

	app.Post("/task", func(ctx iris.Context) {
		var task dto.Task
		ctx.ReadJSON(&task)
		tasksCollection.Insert(&task)
	})

	app.Run(iris.Addr(":8080"))
}
