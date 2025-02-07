package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rajanlagah/go-course/db"
)

// return id of task created on success
func SaveTask(ctx *gin.Context){
	var payload db.PostTaskPayload
	if  err := ctx.ShouldBindJSON(&payload);err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Unable to read the body"})
		return
	}
	id, err := db.TaskRepository.SaveTaskQuery(payload)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": true, "msg": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"error": false, "msg": id})
}

// return tasks on success
func ReadTask(ctx *gin.Context){
	tasks, err := db.TaskRepository.ReadTaskQuery()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": true, "msg": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"error": false, "data": tasks})
}

func UpdateTask(ctx *gin.Context){
	// read payload
	var payload db.UpdateTaskPayload

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// read existing task by payload.ID
	task, err := db.TaskRepository.GetTaskById(payload.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError,  gin.H{"error": true, "msg": err.Error()})
		return
	}

	// override existing task by payload
	
	if payload.Title == "" {
		payload.Title = task.Title
	}

	if payload.Content == "" {
		payload.Content = task.Content
	}

	if payload.Status == "" {
		payload.Status = task.Status
	}

	updateDataErr := db.TaskRepository.UpdateTask(payload)

	if updateDataErr != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": true, "msg": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"error": false,
		"data": payload,
	})
}

func DeleteTask(c *gin.Context){
	taskId := c.Param("id")

	id, err := strconv.Atoi(taskId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "msg": "Invalid Id"})
	}

	deleteErr := db.TaskRepository.DeleteTaskQuery(id)

	if deleteErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": true, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"error": false, "msg": "Task with ID " + taskId + " deleted successfully"})
}