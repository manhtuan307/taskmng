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

	taskAPI.Post("/search/{pageSize:int}/{pageIndex:int}", func(ctx iris.Context) {
		searchTask(ctx, tasksCollection)
	})

	taskAPI.Post("", func(ctx iris.Context) {
		addTask(ctx, tasksCollection)
	})

	taskAPI.Delete("/{taskId:string}", func(ctx iris.Context) {
		deleteTask(ctx, tasksCollection)
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

func searchTask(ctx iris.Context, tasksCollection *mgo.Collection) {
	var searchRequest dto.SearchTaskRequest
	ctx.ReadJSON(&searchRequest)
	log.Print("SearchCondition: ", searchRequest.SearchCondition)

	var searchCondition bson.M
	if len(searchRequest.SearchCondition) > 0 {
		var newSearchCondition = "^.*" + searchRequest.SearchCondition + ".*$"
		searchCondition = bson.M{"name": bson.RegEx{Pattern: newSearchCondition, Options: "i"}}
	} else {
		searchCondition = bson.M{}
	}

	var tasks []dto.Task

	var pageSize, _ = ctx.Params().GetInt("pageSize")
	if pageSize < 1 {
		pageSize = PageSize
	}
	var pageIndex, _ = ctx.Params().GetInt("pageIndex")
	if pageIndex < 1 {
		pageIndex = 1
	}
	var tasksCount, err = tasksCollection.Find(searchCondition).Count()
	var quotient, remainder = utils.DivMod(tasksCount, pageSize)
	var numOfPages = quotient
	if remainder != 0 {
		numOfPages++
	}
	var getTaskResult dto.SearchTasksActionResult
	if numOfPages >= 1 {
		if pageIndex > numOfPages {
			pageIndex = numOfPages
		}
		var numSkip = (pageIndex - 1) * pageSize
		err = tasksCollection.Find(searchCondition).Limit(pageSize).Skip(numSkip).All(&tasks)

		if err != nil {
			log.Fatal(err)
			getTaskResult = dto.SearchTasksActionResult{IsSuccess: false, Message: err.Error()}
		} else {
			getTaskResult = dto.SearchTasksActionResult{IsSuccess: true, Message: "",
				Tasks: tasks, NumOfPages: numOfPages, PageIndex: pageIndex}
		}
	} else {
		getTaskResult = dto.SearchTasksActionResult{IsSuccess: true, Message: "No task found", Tasks: []dto.Task{}}
	}
	ctx.JSON(getTaskResult)
}

func addTask(ctx iris.Context, tasksCollection *mgo.Collection) {
	var task dto.Task
	ctx.ReadJSON(&task)
	task.ID = bson.NewObjectId()
	log.Print("Adding task: ", task)
	tasksCollection.Insert(&task)
}

func deleteTask(ctx iris.Context, tasksCollection *mgo.Collection) {
	var taskID = ctx.Params().Get("taskId")
	if len(taskID) > 0 {
		log.Print("Deleting task: ", taskID)
		err := tasksCollection.Remove(bson.M{"_id": bson.ObjectIdHex(taskID)})
		if err != nil {
			log.Print("Deleting fail: ", err)
		}
	}
}
