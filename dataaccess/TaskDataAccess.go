package dataaccess

import (
	"log"
	"taskmng/dto"
	"taskmng/utils"
	"time"

	"gopkg.in/mgo.v2/bson"
)

//UpdateTask - update a task
func UpdateTask(updatedTask dto.Task, taskID string) {
	if len(taskID) > 0 {
		log.Print("Updating task: ", taskID)
		var query = bson.M{"_id": bson.ObjectIdHex(taskID)}
		var change = bson.M{"$set": bson.M{"name": updatedTask.Name, "status": updatedTask.Status, "updated": time.Now()}}
		err := tasksCollection.Update(query, change)
		if err != nil {
			log.Print("Error: ", err)
			panic(err)
		}
	}
}

//GetTask - get task by its Id
func GetTask(taskID string) dto.Task {
	var task dto.Task
	if len(taskID) > 0 {
		log.Print("Fetching task: ", taskID)
		err := tasksCollection.Find(bson.M{"_id": bson.ObjectIdHex(taskID)}).One(&task)
		if err != nil {
			log.Print("Error: ", err)
			panic(err)
		}
	}
	return task
}

//DeleteTask - delete task by its Id
func DeleteTask(taskID string) {
	if len(taskID) > 0 {
		log.Print("Deleting task: ", taskID)
		err := tasksCollection.Remove(bson.M{"_id": bson.ObjectIdHex(taskID)})
		if err != nil {
			log.Print("Deleting fail: ", err)
			panic(err)
		}
	}
}

//AddTask - add task to DB
func AddTask(task dto.Task) {
	log.Print("Adding task: ", task)
	err := tasksCollection.Insert(&task)
	if err != nil {
		log.Print("Adding Fail: ", err)
		panic(err)
	}
}

//SearchTasks - search tasks by name, using paging
func SearchTasks(searchKeyword string, pageSize int, orgPageIndex int) ([]dto.Task, int, int) {
	var searchCondition bson.M
	var tasks []dto.Task
	var pageIndex = orgPageIndex
	if len(searchKeyword) > 0 {
		var newSearchCondition = "^.*" + searchKeyword + ".*$"
		searchCondition = bson.M{"name": bson.RegEx{Pattern: newSearchCondition, Options: "i"}}
	} else {
		searchCondition = bson.M{}
	}

	var tasksCount, err = tasksCollection.Find(searchCondition).Count()
	var quotient, remainder = utils.DivMod(tasksCount, pageSize)
	var numOfPages = quotient
	if remainder != 0 {
		numOfPages++
	}
	if numOfPages >= 1 {
		if pageIndex > numOfPages {
			pageIndex = numOfPages
		}
		var numSkip = (pageIndex - 1) * pageSize
		err = tasksCollection.Find(searchCondition).Limit(pageSize).Skip(numSkip).All(&tasks)
		if err != nil {
			log.Fatal(err)
			panic(err)
		}
	}

	return tasks, numOfPages, pageIndex
}
