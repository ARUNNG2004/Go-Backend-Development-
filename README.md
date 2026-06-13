# go-user-api

A RESTful User Management API built with **Go**, **Fiber v2**, **MySQL**, and **sqlc**. Features structured logging via Zap, request validation, and a clean layered architecture.

---

## 📦 Tech Stack

| Tool | Purpose |
|------|---------|
| [Go](https://go.dev/) | Programming language |
| [Fiber v2](https://gofiber.io/) | HTTP web framework |
| [MySQL](https://www.mysql.com/) | Relational database |
| [sqlc](https://sqlc.dev/) | Type-safe SQL code generation |
| [go-playground/validator](https://github.com/go-playground/validator) | Request payload validation |
| [Zap](https://github.com/uber-go/zap) | Structured, high-performance logging |
| [godotenv](https://github.com/joho/godotenv) | `.env` file loading |

---

## 📁 Project Structure

```
go-user-api/
├── cmd/
│   └── server/
│       └── main.go          # Application entry point
├── config/                  # Config loader (reads .env)
├── db/
│   ├── migrations/
│   │   └── schema.sql       # Database schema
│   └── sqlc/                # sqlc-generated query code
├── internal/
│   ├── handler/             # HTTP request handlers
│   ├── logger/              # Zap logger initializer
│   ├── models/              # Domain models
│   └── routes/              # Route registration
├── .env                     # Environment variables (not committed)
├── go.mod
├── go.sum
└── sqlc.yaml                # sqlc configuration
```

---

## 🚀 Getting Started

### Prerequisites

- [Go 1.21+](https://go.dev/dl/)
- [MySQL](https://dev.mysql.com/downloads/)
- [sqlc](https://docs.sqlc.dev/en/latest/overview/install.html) *(for regenerating queries)*

### 1. Clone the Repository

```bash
git clone https://github.com/your-username/go-user-api.git
cd go-user-api
```

### 2. Set Up Environment Variables

Copy the example and fill in your values:

```bash
cp .env.example .env
```

`.env` format:

```env
DB_DRIVER=mysql
DB_SOURCE=user:password@tcp(localhost:3306)/dbname?parseTime=true
PORT=:8080
```

### 3. Set Up the Database

Run the migration SQL against your MySQL instance:

```bash
sqlc generate
```

### 4. Install Dependencies

```bash
go mod tidy
```

### 5. Run the Server

```bash
go run cmd/server/main.go
```

The server will start on the port defined in your `.env` (default `:8080`).

---

## 🗄️ Database Schema

```sql
CREATE TABLE users (
    id   BIGINT AUTO_INCREMENT PRIMARY KEY,
    name TEXT   NOT NULL,
    dob  DATE   NOT NULL
);
```

---

## 🔌 API Endpoints

**Base URL:** `https://go-backend-development.onrender.com`

| Method | Endpoint      | Full Render URL | Description          |
|--------|--------------|-----------------|----------------------|
| POST   | `/users`     | `https://go-backend-development.onrender.com/users` | Create a new user    |
| GET    | `/users`     | `https://go-backend-development.onrender.com/users` | List all users       |
| GET    | `/users/:id` | `https://go-backend-development.onrender.com/users/:id` | Get a user by ID     |
| PUT    | `/users/:id` | `https://go-backend-development.onrender.com/users/:id` | Update a user by ID  |
| DELETE | `/users/:id` | `https://go-backend-development.onrender.com/users/:id` | Delete a user by ID  |

### Example: Create User

**Request**

```http
POST https://go-backend-development.onrender.com/users
Content-Type: application/json

{
  "name": "John Doe",
  "dob": "1995-08-15"
}
```

**Response** `201 Created`

```json
{
  "id": 1,
  "name": "John Doe",
  "dob": "1995-08-15"
}
```

### Example: List All Users

**Request**

```http
GET https://go-backend-development.onrender.com/users
```

**Response** `200 OK`

```json
[
  {
    "id": 1,
    "name": "John Doe",
    "dob": "1995-08-15",
    "age": 30
  }
]
```

### Example: Get User by ID

**Request**

```http
GET https://go-backend-development.onrender.com/users/1
```

**Response** `200 OK`

```json
{
  "id": 1,
  "name": "John Doe",
  "dob": "1995-08-15",
  "age": 30
}
```

### Example: Update User by ID

**Request**

```http
PUT https://go-backend-development.onrender.com/users/1
Content-Type: application/json

{
  "name": "John Smith",
  "dob": "1995-08-15"
}
```

**Response** `200 OK`

```json
{
  "id": 1,
  "name": "John Smith",
  "dob": "1995-08-15"
}
```

### Example: Delete User by ID

**Request**

```http
DELETE https://go-backend-development.onrender.com/users/1
```

**Response** `204 No Content`
```

---

## ♻️ Regenerating sqlc Code

After modifying SQL queries in `db/`, regenerate the type-safe Go code:

```bash
sqlc generate
```

---

## 📝 License

This project is licensed under the [MIT License](LICENSE).
