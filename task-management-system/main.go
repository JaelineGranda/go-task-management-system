package main

import (
	"task-management-system/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/tasks", handlers.CreateTask)
	router.GET("/tasks/:id", handlers.GetTask)
	router.PUT("/tasks/:id", handlers.UpdateTask)
	router.DELETE("/tasks/:id", handlers.DeleteTask)
	router.GET("/tasks", handlers.ListTasks)

	router.Run(":8080")
}
