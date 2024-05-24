package storage

import (
	"errors"
	"sync"
	"task-management-system/models"
)

var (
	tasks  = make(map[int]models.Task)
	nextID = 1
	mu     sync.Mutex
)

func CreateTask(task models.Task) models.Task {
	mu.Lock()
	defer mu.Unlock()
	task.ID = nextID
	task.Status = "Pending"
	tasks[nextID] = task
	nextID++
	return task
}

func ReadTask(id int) (models.Task, error) {
	mu.Lock()
	defer mu.Unlock()
	task, exists := tasks[id]
	if !exists {
		return models.Task{}, errors.New("task not found")
	}
	return task, nil
}

func UpdateTask(id int, updatedTask models.Task) (models.Task, error) {
	mu.Lock()
	defer mu.Unlock()
	task, exists := tasks[id]
	if !exists {
		return models.Task{}, errors.New("task not found")
	}
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

func DeleteTask(id int) error {
	mu.Lock()
	defer mu.Unlock()
	_, exists := tasks[id]
	if !exists {
		return errors.New("task not found")
	}
	delete(tasks, id)
	return nil
}

func ListTasks() []models.Task {
	mu.Lock()
	defer mu.Unlock()
	taskList := make([]models.Task, 0, len(tasks))
	for _, task := range tasks {
		taskList = append(taskList, task)
	}
	return taskList
}
