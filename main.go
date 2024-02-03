// main.go
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shubham/taskmanagement/database"
	"github.com/shubham/taskmanagement/routes"
)

func main() {
	db := database.InitDB()
	defer db.Close()

	r := gin.Default()

	// Initialize routes
	routes.InitializeRoutes(r, db)

	// Run the server
	r.Run(":8080")
}
