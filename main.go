package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rajanlagah/go-course/config"
)

func main(){
	handler := gin.Default()
	handler.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Ok from GIN",
		})
	})
	config.Config.LoadConfig()

	server := &http.Server{
		Addr : config.Config.AppPort,
		Handler: handler,
	}
	server.ListenAndServe()
}