package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"task-management-system/handlers"
	"task-management-system/models"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/tasks", handlers.CreateTask)
	router.GET("/tasks/:id", handlers.GetTask)
	router.PUT("/tasks/:id", handlers.UpdateTask)
	router.DELETE("/tasks/:id", handlers.DeleteTask)
	router.GET("/tasks", handlers.ListTasks)
	return router
}

func TestCreateTask(t *testing.T) {
	r := SetUpRouter()
	task := models.Task{Title: "Test Task", Description: "Create test task"}
	jsonValue, _ := json.Marshal(task)
	req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestReadTask(t *testing.T) {
	r := SetUpRouter()
	req, _ := http.NewRequest("GET", "/tasks/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUpdateTask(t *testing.T) {
	r := SetUpRouter()
	task := models.Task{Title: "Updated Task", Description: "Updated description", Status: "In Progress"}
	jsonValue, _ := json.Marshal(task)
	req, _ := http.NewRequest("PUT", "/tasks/1", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeleteTask(t *testing.T) {
	r := SetUpRouter()
	req, _ := http.NewRequest("DELETE", "/tasks/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestListTasks(t *testing.T) {
	r := SetUpRouter()
	req, _ := http.NewRequest("GET", "/tasks", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}
