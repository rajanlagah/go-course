package main

import (
	"context"
	"net/http"

	"github.com/rajanlagah/go-course/config"
	"github.com/rajanlagah/go-course/db"
	"github.com/rajanlagah/go-course/routes"
)

func main(){
	db.InitDB()
	handler := routes.MounteRoutes()
	// handler := gin.Default()
	// handler.GET("/", func(ctx *gin.Context) {
	// 	ctx.JSON(http.StatusOK, gin.H{
	// 		"message": "Ok from AIR",
	// 	})
	// })

	server := &http.Server{
		Addr : config.Config.AppPort,
		Handler: handler,
	}
	defer db.DB.Close(context.Background())
	server.ListenAndServe()
}