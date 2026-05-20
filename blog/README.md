# Blog Backend API

Simple blog backend project built using Go and PostgreSQL.

This project implements basic CRUD operations for blogs using:

- Go
- net/http
- PostgreSQL
- pgx

---

# Project Structure

```text
blog/
├── cmd/
│   └── server/
│       └── main.go
│
├── internal/
│   ├── handler/
│   ├── model/
│   ├── repository/
│   ├── routes/
│   └── service/
│
├── pkg/
│   ├── config/
│   └── db/
│
├── migrations/
│   ├── schema.up.sql
│   └── schema.down.sql
│
├── .env
├── go.mod
└── README.md
```

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
| POST | `/blog` | Create blog |
| GET | `/blogs` | Get all blogs |
| PUT | `/blogs/update?id=1` | Update blog |
| DELETE | `/blogs/delete?id=1` | Delete blog |

---

# Architecture

This project follows layered architecture:

```text
Routes Layer
    ↓
Handler Layer
    ↓
Service Layer
    ↓
Repository Layer
    ↓
DB Layer
    ↓
PostgreSQL
```

---
