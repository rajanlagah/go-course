package handlers

import (
	"net/http"

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