package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rajanlagah/go-course/db"
)

type PostTaskPayload struct {
	Title string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Status string `json:"status"`
}

func SaveTask(ctx *gin.Context){
	var payload PostTaskPayload
	if err := ctx.ShouldBindJSON(&payload); err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Unable to read the body"})
		return
	}
	var id int
	query := `INSERT into tasks (title, description, status) VALUES ($1,$2,$3) RETURNING id;`
	log.Print(db.DB)
	log.Print("db.DB")
	err := db.DB.QueryRow(context.Background(), query, payload.Title, payload.Description,payload.Status).Scan(&id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": true, "msg": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"error": false, "msg": "id"})
}