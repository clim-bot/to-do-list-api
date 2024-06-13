# To-Do List API

This is a simple RESTful API for managing a to-do list, built with Golang using the Gin framework, Gorm for ORM, PostgreSQL for the database, and JWT for authentication.

## Features

- User Registration and Login
- JWT-based Authentication
- Add, Update, Delete, and Retrieve Tasks
- Unit Tests
- Docker for Containerization

## Tech Stack

- Gin (for routing)
- Gorm (for ORM)
- PostgreSQL (for database)
- JWT (for authentication)
- Docker (for containerization)

## Project Structure
```
to-do-list-api/
├── Dockerfile
├── docker-compose.yml
├── go.mod
├── go.sum
├── main.go
├── controllers
│ ├── auth_controller.go
│ └── task_controller.go
├── config
│ ├── db.go
├── models
│ ├── task.go
│ └── user.go
├── repositories
│ ├── task_repository.go
│ └── user_repository.go
├── routes
│ └── routes.go
├── tests
│ ├── auth_controller_test.go
│ └── task_controller_test.go
└── utils
├── auth.go
└── hash.go
```

## Getting Started

### Prerequisites

- Docker
- Docker Compose

### Setup

1. **Clone the repository**

   ```sh
   git clone https://github.com/yourusername/to-do-list-api.git
   cd to-do-list-api

2. Create a .env file
    ```sh
    touch .env
    ```

    Add the following content to the .env file:
    ```sh
    DSN_URL=postgres://postgres:postgres@db:5432/todo_db?sslmode=disable

    ```
    The API will be available at http://localhost:8080.


### API Endpoints
#### User Registration
- URL: /auth/register
- Method: POST
- Data: {"email": "user@example.com", "password": "password123"}
```sh
curl -X POST http://localhost:8080/auth/register \
-H "Content-Type: application/json" \
-d '{"email": "user@example.com", "password": "password123"}'
```

#### User Login
- URL: /auth/login
- Method: POST
- Data: {"email": "user@example.com", "password": "password123"}
```sh
curl -X POST http://localhost:8080/auth/login \
-H "Content-Type: application/json" \
-d '{"email": "user@example.com", "password": "password123"}'

```

#### Create a Task
- URL: /tasks
- Method: POST
- Headers: Authorization: <JWT_TOKEN>
- Data: {"title": "New Task", "status": "incomplete"}
```sh
curl -X POST http://localhost:8080/tasks \
-H "Content-Type: application/json" \
-H "Authorization: <JWT_TOKEN>" \
-d '{"title": "New Task", "status": "incomplete"}'
```

#### Get All Tasks
- URL: /tasks
- Method: GET
- Headers: Authorization: <JWT_TOKEN>
```sh
curl -X GET http://localhost:8080/tasks \
-H "Content-Type: application/json" \
-H "Authorization: <JWT_TOKEN>"
```

#### Update a Task
- URL: /tasks
- Method: PUT
- Headers: Authorization: <JWT_TOKEN>
- Data: {"id": 1, "title": "Updated Task", "status": "complete"}
```sh
curl -X PUT http://localhost:8080/tasks \
-H "Content-Type: application/json" \
-H "Authorization: <JWT_TOKEN>" \
-d '{"id": 1, "title": "Updated Task", "status": "complete"}'
```

#### Delete a Task
- URL: /tasks
- Method: DELETE
- Headers: Authorization: <JWT_TOKEN>
- Data: {"id": 1}
```sh
curl -X DELETE http://localhost:8080/tasks \
-H "Content-Type: application/json" \
-H "Authorization: <JWT_TOKEN>" \
-d '{"id": 1}'
```

### Running Tests
To run the tests, use the following command:
```sh
go test ./tests/...
```

### License
This project is licensed under the MIT License.