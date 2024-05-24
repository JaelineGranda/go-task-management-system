package handlers

import (
	"net/http"
	"strconv"
	"task-management-system/storage"

	"github.com/gin-gonic/gin"
)

// Delete a task by its ID
func DeleteTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	// Call storage layer to delete the task
	if err := storage.DeleteTask(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	// Return a success message upon successful deletion
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}
