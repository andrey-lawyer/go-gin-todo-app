# Todo App

A simple Todo application built with Go, Gin, and MongoDB.

## About the Project

This project is a basic task management (todo) application that allows users to register, log in, and manage their personal tasks. The backend is written in Go using the Gin web framework, and MongoDB is used as the database. User authentication is session-based (cookies).

## Technologies Used

- **Go** (Golang)
- **Gin** (HTTP web framework)
- **MongoDB** (NoSQL database)
- **Docker** & **Docker Compose** (for running MongoDB)
- **Sessions** (via cookies)
- **bcrypt** (for password hashing)

## Project Structure

```
.
├── config/           # Configuration and environment loading
├── errors/           # Custom error types
├── handlers/         # HTTP handlers (controllers)
├── infrastructure/   # Database connection
├── middleware/       # Custom middleware (auth, error recovery)
├── models/           # Data models (User, Task)
├── repositories/     # Data access layer (MongoDB)
├── services/         # Business logic
├── main.go           # Application entry point
├── docker-compose.yml
└── README.md
```

## API Endpoints

| Method | Endpoint      | Description           | Auth Required |
|--------|--------------|----------------------|--------------|
| POST   | /register    | Register new user    | No           |
| POST   | /login       | Login user           | No           |
| GET    | /tasks       | Get user's tasks     | Yes          |
| POST   | /tasks       | Create new task      | Yes          |

### Example Request/Response

**Register**
```http
POST /register
Content-Type: application/json

{
  "username": "john",
  "password": "Secret123"
}
```

**Login**
```http
POST /login
Content-Type: application/json

{
  "username": "john",
  "password": "Secret123"
}
```

**Get Tasks**
```http
GET /tasks
Cookie: auth-session=...
```

**Create Task**
```http
POST /tasks
Content-Type: application/json
Cookie: auth-session=...

{
  "title": "Buy milk",
  "description": "2 liters"
}
```

## Getting Started

### Prerequisites

- [Go](https://golang.org/dl/) 1.18 or newer
- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/)

### 1. Clone the repository

```bash
git clone https://github.com/your-username/todo-app.git
cd todo-app
```

### 2. Start MongoDB with Docker Compose

```bash
docker-compose up -d
```

This will start a MongoDB instance on port 27017.

### 3. Configure environment variables

Create a `.env` file in the project root (optional, defaults are provided):

```
MONGO_URI=mongodb://localhost:27017
PORT=8080
SESSION_SECRET=your_secret_key
```

### 4. Run the application locally

```bash
go run main.go
```

The server will start on `http://localhost:8080`.

## Notes

- MongoDB runs in Docker, the Go application runs locally.
- All data is stored in the `todo_db` database in MongoDB.
- Sessions are stored in cookies.


