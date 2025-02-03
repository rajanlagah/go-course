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

	server := &http.Server{
		Addr : config.Config.AppPort,
		Handler: handler,
	}
	server.ListenAndServe()
	defer db.DB.Close(context.Background())
}