package main

import (
	"github.com/gin-gonic/gin"
	"github.com/krishna/go-tasker/internal/models"
	"github.com/krishna/go-tasker/internal/routes"

	"log"
)

func main() {
	// Initialize Gin router
	router := gin.Default()

	// Connect to the database
	models.ConnectDatabase()

	// Register routes
	routes.RegisterRoutes(router)

	// Run the server
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

// package main

// import (
// 	// _ "go-tasker/docs"
// 	"log"

// 	"github.com/gin-gonic/gin"
// 	"github.com/krishna/go-tasker/internal/models"
// 	"github.com/krishna/go-tasker/internal/routes"
// 	// "github.com/swaggo/gin-swagger"
// )

// // @title GoTasker API
// // @version 1.0
// // @description This is a sample server for GoTasker.
// // @host localhost:8080
// // @BasePath /

// func main() {
// 	// Initialize Gin router
// 	router := gin.Default()

// 	// Connect to the database
// 	models.ConnectDatabase()

// 	// Register routes
// 	routes.RegisterRoutes(router)

// 	// Swagger endpoint
// 	// router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

// 	// Run the server
// 	if err := router.Run(":8080"); err != nil {
// 		log.Fatalf("Server failed to start: %v", err)
// 	}
// }
