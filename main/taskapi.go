package main

import (
	"log"
	"taskmng/dataaccess"
	"taskmng/dto"

	"gopkg.in/mgo.v2/bson"

	"github.com/kataras/iris"
)

func searchTask(ctx iris.Context) {
	var searchRequest dto.SearchTaskRequest
	ctx.ReadJSON(&searchRequest)
	log.Print("Search Task by condition: ", searchRequest.SearchCondition)
	var pageSize, _ = ctx.Params().GetInt("pageSize")
	if pageSize < 1 {
		pageSize = PageSize
	}
	var paramPageIndex, _ = ctx.Params().GetInt("pageIndex")
	if paramPageIndex < 1 {
		paramPageIndex = 1
	}
	var tasks, numOfPages, pageIndex = dataaccess.SearchTasks(searchRequest.SearchCondition, pageSize, paramPageIndex)
	var getTaskResult dto.SearchTasksActionResult
	if numOfPages >= 1 {
		getTaskResult = dto.SearchTasksActionResult{IsSuccess: true, Message: "",
			Tasks: tasks, NumOfPages: numOfPages, PageIndex: pageIndex}
	} else {
		getTaskResult = dto.SearchTasksActionResult{IsSuccess: true, Message: "No task found", Tasks: []dto.Task{}}
	}
	ctx.JSON(getTaskResult)
}

func addTask(ctx iris.Context) {
	var task dto.Task
	ctx.ReadJSON(&task)
	var userID = ctx.Values().GetString("UserID")
	task.UserID = bson.ObjectIdHex(userID)
	dataaccess.AddTask(task)
	var result = dto.ActionResult{IsSuccess: true, Message: ""}
	ctx.JSON(result)
}

func deleteTask(ctx iris.Context) {
	var taskID = ctx.Params().Get("taskId")
	var userID = ctx.Values().GetString("UserID")
	if len(taskID) > 0 {
		dataaccess.DeleteTask(taskID, userID)
	}
}

func getTask(ctx iris.Context) {
	var taskID = ctx.Params().Get("taskId")
	var task = dataaccess.GetTask(taskID)
	ctx.JSON(task)
}

func updateTask(ctx iris.Context) {
	var updatedTask dto.Task
	ctx.ReadJSON(&updatedTask)
	var userID = ctx.Values().GetString("UserID")
	var taskID = ctx.Params().Get("taskId")
	if len(taskID) > 0 {
		dataaccess.UpdateTask(updatedTask, taskID, userID)
	}
	var result = dto.ActionResult{IsSuccess: true, Message: ""}
	ctx.JSON(result)
}
