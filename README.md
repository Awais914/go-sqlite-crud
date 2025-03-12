# Go Students API

A simple RESTful API built with Go for managing student records.

## Features

- CRUD operations for student records
- SQLite database storage
- Graceful server shutdown
- Request validation
- Pagination support for listing students

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/api/students` | Create a new student |
| GET | `/api/students/{id}` | Get a student by ID |
| GET | `/api/students` | List all students (with pagination) |
| PUT | `/api/students/{id}` | Update a student by ID |
| DELETE | `/api/students/{id}` | Delete a student by ID |

## Getting Started

### Prerequisites

- Go 1.22 or higher
- SQLite

### Installation

1. Clone the repository

```bash
git clone https://github.com/yourusername/go-students-api.git
cd go-students-api
```

2. Install dependencies

```bash
go mod tidy
```

3. Run the server

```bash
go run cmd/go-student-api/main.go
```