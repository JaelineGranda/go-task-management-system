package storage

import (
	"errors"
	"sync"
	"task-management-system/models"
)

// Global variables to store tasks and manage task IDs
var (
	tasks  = make(map[int]models.Task) // In-memory storage for tasks
	nextID = 1
	mu     sync.Mutex
)

// Add a new task to the storage and return created task
func CreateTask(task models.Task) models.Task {
	mu.Lock()
	defer mu.Unlock()
	// Assign ID and initial default status to the new task
	task.ID = nextID
	task.Status = "Pending"
	tasks[nextID] = task
	nextID++
	return task
}

// Get a task by its ID
func ReadTask(id int) (models.Task, error) {
	mu.Lock()
	defer mu.Unlock()
	// Check if task exists
	task, exists := tasks[id]
	if !exists {
		return models.Task{}, errors.New("task not found")
	}
	return task, nil
}

// UUpdate an existing task
func UpdateTask(id int, updatedTask models.Task) (models.Task, error) {
	mu.Lock()
	defer mu.Unlock()
	// Check if task exists
	task, exists := tasks[id]
	if !exists {
		return models.Task{}, errors.New("task not found")
	}
	// Update the task fields
	if updatedTask.Title != "" {
		task.Title = updatedTask.Title
	}
	if updatedTask.Description != "" {
		task.Description = updatedTask.Description
	}
	if updatedTask.Status != "" {
		task.Status = updatedTask.Status
	}
	tasks[id] = task
	return task, nil
}

// Removes a task from the storage by its ID
func DeleteTask(id int) error {
	mu.Lock()
	defer mu.Unlock()
	// Check if the task exists
	_, exists := tasks[id]
	if !exists {
		return errors.New("task not found")
	}
	delete(tasks, id)
	return nil
}

// Returns a list of all tasks
func ListTasks() []models.Task {
	mu.Lock()
	defer mu.Unlock()
	// Create a list of all tasks
	taskList := make([]models.Task, 0, len(tasks))
	for _, task := range tasks {
		taskList = append(taskList, task)
	}
	return taskList
}
