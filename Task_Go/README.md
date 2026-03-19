# Task Go API

A simple Task Manager backend built using **Go**, **Fiber**, and **PostgreSQL** with **JWT Authentication**.  
This project provides CRUD APIs to manage tasks with protected routes.

---

## Features

- User Registration and Login  
- JWT-based Authentication  
- Protected Task APIs  
- Create a task  
- Get all tasks  
- Update a task  
- Delete a task  
- PostgreSQL integration using GORM  
- Environment variable configuration  
- Basic logging using Zerolog  

---

## Tech Stack

- Go (Golang)  
- Fiber (Web Framework)  
- GORM (ORM)  
- PostgreSQL  
- JWT (Authentication)  
- Zerolog (Logging)  

---

## Project Structure

task_go/

├── config/        # Config and DB connection  
├── handlers/      # API logic (auth + tasks)  
├── middleware/    # JWT middleware  
├── models/        # Data models  
├── routes/        # API routes  
├── utils/         # JWT and logger  
├── .env           # Environment variables  
└── main.go        # Entry point  

---

## API Endpoints

### Auth APIs

| Method | Endpoint   | Description        |
|--------|-----------|--------------------|
| POST   | /register | Register user      |
| POST   | /login    | Login user & get token |

---

### Task APIs (Protected)

| Method | Endpoint     | Description   |
|--------|-------------|--------------|
| GET    | /tasks      | Get all tasks |
| POST   | /tasks      | Create task   |
| PUT    | /tasks/:id  | Update task   |
| DELETE | /tasks/:id  | Delete task   |

---

## Authentication

- After login, a JWT token is generated  
- Token must be sent in headers for protected routes  

Example:

Authorization: Bearer your_token_here  

---

## Run Locally

1. Clone the repository  
git clone <your-repo-url>  
cd task_go  

2. Install dependencies  
go mod tidy  

3. Create `.env` file  

Example:

DB_DSN=host=localhost user=your_user password=your_password dbname=task_go port=5432 sslmode=disable  
JWT_SECRET=your_secret_key  

4. Run the server  
go run main.go  

Server will run on:  
http://localhost:4000  

---

## Example Requests

### Register User

POST /register  

Body:
{
  "username": "testuser",
  "password": "1234"
}

---

### Login User

POST /login  

Body:
{
  "username": "testuser",
  "password": "1234"
}

---

### Create Task (Protected)

POST /tasks  

Headers:
Authorization: Bearer your_token  

Body:
{
  "title": "Learn Fiber",
  "completed": false
}

---

## Notes

- This project is built for learning backend development in Go  
- Code is kept simple for better understanding  
- Not production-ready (no password hashing, minimal validation)  

---

## Future Improvements

- Add password hashing (bcrypt)  
- Add request validation  
- Add user-task relationship  
- Implement refresh token system  
- Improve error handling  