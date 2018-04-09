package main

import (
	"log"
	"taskmng/dto"
	"taskmng/utils"

	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/middleware/logger"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// PageSize is default number of records for paging
const PageSize = 20

func main() {
	app := iris.New()
	app.Use(logger.New())

	// Database connection
	session, err := mgo.Dial("localhost:27017")
	if nil != err {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	tasksCollection := session.DB("task_management").C("Task")

	appCors := configCors()
	var taskAPI = app.Party("/task", appCors).AllowMethods(iris.MethodOptions)

	taskAPI.Get("/{pageSize:int}/{pageIndex:int}", func(ctx iris.Context) {
		getTask(ctx, tasksCollection)
	})

	taskAPI.Post("", func(ctx iris.Context) {
		addTask(ctx, tasksCollection)
	})

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

func getTask(ctx iris.Context, tasksCollection *mgo.Collection) {
	var tasks []dto.Task

	var pageSize, _ = ctx.Params().GetInt("pageSize")
	if pageSize < 1 {
		pageSize = PageSize
	}
	var pageIndex, _ = ctx.Params().GetInt("pageIndex")
	if pageIndex < 1 {
		pageIndex = 1
	}
	var tasksCount, err = tasksCollection.Find(bson.M{}).Count()
	var quotient, remainder = utils.DivMod(tasksCount, pageSize)
	var numOfPages = quotient
	if remainder != 0 {
		numOfPages++
	}
	var getTaskResult dto.GetTasksActionResult
	if numOfPages >= 1 {
		if pageIndex > numOfPages {
			pageIndex = numOfPages
		}
		var numSkip = (pageIndex - 1) * pageSize

		err = tasksCollection.Find(bson.M{}).Limit(pageSize).Skip(numSkip).All(&tasks)

		if err != nil {
			log.Fatal(err)
			getTaskResult = dto.GetTasksActionResult{IsSuccess: false, Message: err.Error()}
		} else {
			getTaskResult = dto.GetTasksActionResult{IsSuccess: true, Message: "",
				Tasks: tasks, NumOfPages: numOfPages, PageIndex: pageIndex}
		}
	} else {
		getTaskResult = dto.GetTasksActionResult{IsSuccess: true, Message: "No task found", Tasks: []dto.Task{}}
	}
	ctx.JSON(getTaskResult)
}

func addTask(ctx iris.Context, tasksCollection *mgo.Collection) {
	var task dto.Task
	ctx.ReadJSON(&task)
	tasksCollection.Insert(&task)
}
