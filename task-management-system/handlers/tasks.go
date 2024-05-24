package handlers

import (
	"net/http"
	"strconv"
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

// Fetches a task by its ID
func ReadTask(c *gin.Context) {
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

// List all tasks
func ListTasks(c *gin.Context) {
	// Call storage layer to list all tasks
	tasks := storage.ListTasks()
	// Return the list of tasks as JSON response
	c.JSON(http.StatusOK, tasks)
}
