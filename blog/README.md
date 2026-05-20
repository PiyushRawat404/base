# Blog Backend API

Simple blog backend project built using Go and PostgreSQL.

This project implements basic CRUD operations for blogs using:

- Go
- net/http
- PostgreSQL
- pgx
- layered architecture

---

# Project Structure

```text
blog-backend/
├── cmd/
│   └── main.go
│
├── internal/
│   ├── config/
│   ├── database/
│   ├── handler/
│   ├── model/
│   ├── repository/
│   ├── routes/
│   └── service/
│
├── migrations/
│   ├── 001_create_blogs_table.up.sql
│   └── 001_create_blogs_table.down.sql
│
├── .env
├── go.mod
└── README.md
```

---

# Features

- Create blog
- Get all blogs
- Update blog
- Delete blog

---

# Technologies Used

- Golang
- PostgreSQL
- pgx driver
- net/http

---

# API Endpoints

| Method | Endpoint | Description |
|---|---|---|
| POST | `/blogs/create` | Create blog |
| GET | `/blogs` | Get all blogs |
| PUT | `/blogs/update?id=1` | Update blog |
| DELETE | `/blogs/delete?id=1` | Delete blog |

---



# Architecture

This project follows layered architecture:

```text
Handler Layer
    ↓
Service Layer
    ↓
Repository Layer
    ↓
PostgreSQL
```

---

