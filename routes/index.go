package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rajanlagah/go-course/routes/handlers"
)

func MounteRoutes() *gin.Engine{
	handler := gin.Default()
	handler.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Ok from AIR",
		})
	})

	handler.POST("/task", handlers.SaveTask)
	handler.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"message":"Route not found"})
	})
	return handler
}