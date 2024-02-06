// routes/routes.go
package routes

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/shubham/taskmanagement/handlers"
)

func InitializeRoutes(router *gin.Engine, db *sql.DB) {
	// API endpoints
	router.POST("/tasks", handlers.CreateTask)
	router.GET("/tasks/:id", handlers.GetTask)
	router.PUT("/tasks/:id", handlers.UpdateTask)
	router.DELETE("/tasks/:id", handlers.DeleteTask)
	router.GET("/tasks", handlers.ListTask)
}
