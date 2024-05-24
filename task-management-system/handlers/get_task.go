package handlers

import (
	"net/http"
	"strconv"
	"task-management-system/storage"

	"github.com/gin-gonic/gin"
)

// Fetches a task by its ID
func GetTask(c *gin.Context) {
	// Parse the task ID from URL parameter
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	// Call storage layer to read the task by ID
	task, err := storage.ReadTask(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	// Return the task as JSON response
	c.JSON(http.StatusOK, task)
}

// List all tasks
func ListTasks(c *gin.Context) {
	// Call storage layer to list all tasks
	tasks := storage.ListTasks()
	// Return the list of tasks as JSON response
	c.JSON(http.StatusOK, tasks)
}
