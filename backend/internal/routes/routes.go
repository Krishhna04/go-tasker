package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/krishna/go-tasker/internal/controllers"
	"github.com/krishna/go-tasker/internal/middleware"
)

func RegisterRoutes(router *gin.Engine) {
	// User routes
	userRoutes := router.Group("/user")
	{
		userRoutes.POST("/register", controllers.Register)
		userRoutes.POST("/login", controllers.Login)
	}

	// Task routes with authentication middleware
	taskRoutes := router.Group("/tasks", middleware.AuthMiddleware)
	{
		taskRoutes.GET("/", controllers.GetTasks)
		taskRoutes.POST("/", controllers.CreateTask)
		taskRoutes.GET("/:id", controllers.GetTask)
		taskRoutes.PUT("/:id", controllers.UpdateTask)
		taskRoutes.DELETE("/:id", controllers.DeleteTask)
	}
}
