package handlers

import (
	"net/http"
	"task-management-system/models"
	"task-management-system/storage"

	"github.com/gin-gonic/gin"
)

// Creates a new task
func CreateTask(c *gin.Context) {
	var task models.Task
	// Bind JSON input to task object, checks for required fields
	if err := c.ShouldBindJSON(&task); err != nil || task.Title == "" || task.Description == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	// Call storage layer to create the task, return created task
	createdTask := storage.CreateTask(task)
	c.JSON(http.StatusCreated, createdTask)
}
