package handlers

import (
	"net/http"
	"strconv"
	"task-management-system/models"
	"task-management-system/storage"

	"github.com/gin-gonic/gin"
)

// Updates an existing task
func UpdateTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	var task models.Task
	// Bind JSON input to task object
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	// Call storage layer to update the task and return updated task
	updatedTask, err := storage.UpdateTask(id, task)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedTask)
}
