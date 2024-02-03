// handlers/task_handler.go
package handlers

import (
	"database/sql"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
	Status      string    `json:"status" binding:"oneof=pending in_progress completed"`
}

func CreateTask(c *gin.Context) {

	var newTask Task

	if err := c.BindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Insert new task into the database
	result, err := execDB("INSERT INTO tasks (title, description, due_date, status) VALUES (?, ?, ?, ?)",
		newTask.Title, newTask.Description, newTask.DueDate, newTask.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Get the ID of the inserted task
	taskID, _ := result.LastInsertId()

	// Set the ID in the newTask struct
	newTask.ID = int(taskID)

	c.JSON(http.StatusCreated, newTask)
}

func GetTask(c *gin.Context) {
	taskID := c.Param("id")

	var task Task
	row := db.QueryRow("SELECT id, title, description, due_date, status FROM tasks WHERE id = ?", taskID)
	err := row.Scan(&task.ID, &task.Title, &task.Description, &task.DueDate, &task.Status)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	c.JSON(http.StatusOK, task)
}

func UpdateTask(c *gin.Context) {
	taskID := c.Param("id")
	var updatedTask Task

	if err := c.BindJSON(&updatedTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update the task in the database
	_, err := execDB("UPDATE tasks SET title = ?, description = ?, due_date = ?, status = ? WHERE id = ?",
		updatedTask.Title, updatedTask.Description, updatedTask.DueDate, updatedTask.Status, taskID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	// Set the ID in the updatedTask struct
	updatedTask.ID, _ = strconv.Atoi(taskID)

	c.JSON(http.StatusOK, updatedTask)
}

func DeleteTask(c *gin.Context) {
	taskID := c.Param("id")

	// Delete the task from the database
	_, err := execDB("DELETE FROM tasks WHERE id = ?", taskID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}

func ListTasks(c *gin.Context) {
	// Implementation for listing all tasks
}
