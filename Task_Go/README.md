# Task Go API

A simple Task Manager backend built using **Go**, **Gin**, and **PostgreSQL**.  
This project provides basic CRUD APIs to manage tasks.

---

## Features

- Create a task  
- Get all tasks  
- Update a task  
- Delete a task  
- PostgreSQL database integration using GORM  

---

## Tech Stack

- Go (Golang)  
- Gin (Web Framework)  
- GORM (ORM)  
- PostgreSQL  

---

## Project Structure

task_go/
├── config/     # Database connection  
├── handlers/   # API logic  
├── models/     # Data models  
├── routes/     # API routes  
└── main.go     # Entry point  

---

## API Endpoints

| Method | Endpoint     | Description   |
|--------|-------------|--------------|
| GET    | /tasks      | Get all tasks |
| POST   | /tasks      | Create task   |
| PUT    | /tasks/:id  | Update task   |
| DELETE | /tasks/:id  | Delete task   |

---

## Run Locally

1. Clone the repository  
git clone <your-repo-url>  
cd task_go  

2. Install dependencies  
go mod tidy  

3. Setup PostgreSQL and update DB credentials in config/db.go  

4. Run the server  
go run main.go  

Server will run on:  
http://localhost:4000  

---

## Example Request

POST /tasks  

Body:
{
  "title": "Learn Go",
  "completed": false
}

---

