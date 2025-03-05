package routes

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rajanlagah/go-course/config"
	"github.com/rajanlagah/go-course/routes/handlers"
)

func MounteRoutes() *gin.Engine{
	handler := gin.Default()
	handler.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", config.Config.FEOriginURL},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, 
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"}, 
		ExposeHeaders:    []string{"Content-Length", "Authorization"}, 
		AllowCredentials: true,
		MaxAge: 12 * time.Hour,
	  }))

	handler.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Ok from AIR",
		})
	})

	taskRoutes := handler.Group("/task")
	{
		taskRoutes.PATCH("/", handlers.UpdateTask)
		taskRoutes.POST("/", handlers.SaveTask)
		taskRoutes.GET("/", handlers.ReadTask)
		taskRoutes.DELETE("/:id", handlers.DeleteTask)
	}

	userLoginRoutes := handler.Group("/login")
	{
		userLoginRoutes.GET("/google",  handlers.HandleGoogleLogin)
	}
	callbackLoginRoutes := handler.Group("/callback")
	{
		callbackLoginRoutes.GET("/google",  handlers.HandleGoogleCallback)
	}

	handler.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"message":"Route not found"})
	})
	return handler
}