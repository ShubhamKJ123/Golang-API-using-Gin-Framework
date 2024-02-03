// routes/routes.go
package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/shubham/taskmanagement/database"
)

func InitializeRoutes(router *gin.Engine, db *sql.DB) {
	// API endpoints
	router.POST("/tasks", createTask)
	router.GET("/tasks/:id", getTask)
	router.PUT("/tasks/:id", updateTask)
	router.DELETE("/tasks/:id", deleteTask)
	router.GET("/tasks", listTasks)
}
