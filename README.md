# RESTful API Task Management System

## RESTful API in Go using the gin-gonic/gin framework to manage tasks

Features implemented:
- CRUD endpoints and List endpoint 
- In-memory data storage
- Input validation for creating/updating tasks
- Basic error handling
- Unit tests for each endpoint

## Project structure:

go-task management-system

```
    task-management-system/
        ├── go.mod
        ├── go.sum
        ├── main.go
        ├── handlers/
        │   └── tasks.go
        ├── models/
        │   └── task.go
        ├── storage/
        │   └── memory.go
        └── tests/
            └── tasks_test.go
```

## Set Up Environment

To set up the environment for running the Go RESTful API with the gin-gonic/gin framework, you need to install the following:
1. Go Programming Language: Ensure you have Go installed.
2. Gin Framework: Install the gin-gonic/gin package.
3. Testify: Install the testify package for writing tests.

## Running and Testing the Application:

#### To run the application:

    cd task-management-system


    go run main.go

#### To run the tests:

    go test ./tests/


## Use curl to Test the API:

1. Create a New Task:

   `curl -X POST http://localhost:8080/tasks -H "Content-Type: application/json" -d '{"title": "Task 1", "description": "This is task 1"}'`


3. Get a specific task:

   `curl http://localhost:8080/tasks/1`

6. Get all tasks:

    `curl http://localhost:8080/tasks`

7. Update an existing task:

    `curl -X PUT http://localhost:8080/tasks/1 -H "Content-Type: application/json" -d '{"title": "Updated Task 1", "description": "Updated description", "status": "In Progress"}'`

8. Delete a task:

   `curl -X DELETE http://localhost:8080/tasks/1`



*Optional: Install and use Postman to test the API*

## Example of Using Postman:

- Create a Task
    ```
        Method: POST
        URL: http://localhost:8080/tasks
        Headers: Content-Type: application/json
        Body:
        
        json
        
        {
            "title": "Task 1",
            "description": "This is task 1"
        }
    ```

- Read a Task:
   ```
        Method: GET
        URL: http://localhost:8080/tasks/1
    ```
- Update a Task:
   ```
        Method: PUT
        URL: http://localhost:8080/tasks/1
        Headers: Content-Type: application/json
        Body:

        json

        {
            "title": "Updated Task 1",
            "description": "Updated description",
            "status": "In Progress"
        }
    ```
- Delete a Task:
   ```
        Method: DELETE
        URL: http://localhost:8080/tasks/1
    ```

- List All Tasks:
  ```
        Method: GET
        URL: http://localhost:8080/tasks
  ```


