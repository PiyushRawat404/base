# Practice Project

This repository is created for practicing assignment and revise concepts.

---

## Progress

<details>
<summary><strong>Week 1</strong></summary>

<br>

<details>
<summary><strong>Day 1</strong></summary>

### Git Tasks

- Install Git
- Create GitLab account
- Create repository on GitLab
- Clone repository
- Create branch (`main`, `staging`, `dev`)
- Create feature branch from `dev` branch
- Add `README.md` file
- Commit changes on feature branch
- Push changes to feature branch
- Create a pull request to merge **feature branch → dev**
- Create a pull request to merge **dev → staging**
- Create a pull request to merge **staging → main**

### Project Management Tool

- Create a board on Trello for managing tasks

</details>

<details>
<summary><strong>Day 2</strong></summary>

### HTML CSS JS Tasks

- Make a Task Manager Application
- Add functionality to add tasks
- Add functionality to remove tasks
- Strikethrough the tasks that are completed

</details>

<details>
<summary><strong>Day 3</strong></summary>

### Node.js Backend Integration

- Convert the Task Manager to use Node.js and Express
- Create API routes for tasks
- Store tasks in a JSON file instead of local storage
- Connect frontend with backend using `fetch`
- Add functionality to create, update and delete tasks using API
- Show uncompleted tasks on top and completed tasks at bottom

</details>

<details>
<summary><strong>Day 4</strong></summary>

### Docker Integration

- Install Docker and understand containerization concepts
- Create a `Dockerfile` for the Task Manager application
- Build a Docker image for the application
- Run the application inside a Docker container
- Map container ports to host ports
- Mount a volume to persist JSON file data
- Understand Docker commands such as `docker build`, `docker run`, `docker ps`, `docker stop`
- Learn the difference between **Docker Image** and **Docker Container**
- Ensure the application runs inside Docker with persistent data storage

</details>



</details>

---

<details>
<summary><strong>Week 2</strong></summary>

<br>
<details>
<summary><strong>Day 1</strong></summary>

### Software Testing, SDLC and STLC Practice

- Studied the concepts of Software Development Life Cycle (SDLC) and Software Testing Life Cycle (STLC) to understand how testing is integrated into different phases of software development.
- Explored different types of testing such as functional testing, regression testing, positive testing and negative testing to understand their role in ensuring software quality.
- Created and organized a testing task board on Trello to manage and track testing activities.
- Performed positive and negative testing on real applications such as the Telegram mobile application and the Airbnb website to observe how applications behave with valid and invalid inputs.
- Researched different software testing tools and learned how they are used in both manual and automated testing processes.
- Gained practical understanding of testing concepts by applying testing techniques on real-world applications.

</details>
<details>

<summary><strong>Day 2</strong></summary>

### Go Language Learning and API Development

- Learned the fundamentals of Go (Golang) including basic syntax, structs, slices, maps, and how Go handles project structure.
- Understood how REST APIs are built in Go using the `net/http` package.
- Implemented a simple Go application that exposes REST endpoints.
- Created a **POST API** to add tasks and a **GET API** to retrieve tasks from the server.
- Structured the project using folders such as **handlers, models, routes, and storage** to organize the application code.
- Tested the API endpoints locally to verify the functionality of the GET and POST operations.

</details>
<details>
<summary><strong>Day 3</strong></summary>

### Go Backend with Database and Framework Integration

- Integrated PostgreSQL database with the Go application using GORM ORM for persistent data storage.
- Configured database connection and performed auto-migration for the Task model.
- Implemented full CRUD APIs (Create, Read, Update, Delete) for task management.
- Switched from the default `net/http` package to the Gin framework for better routing and cleaner code structure.
- Created API endpoints for GET, POST, PUT, and DELETE operations using Gin.
- Handled JSON request/response using Gin context methods.
- Tested all API endpoints using Postman to ensure correct functionality.
- Understood how ID auto-generation works in PostgreSQL using sequences.

</details>
<details>
<summary><strong>Day 4</strong></summary>

### Authentication, Middleware and Framework Upgrade

- Replaced Gin framework with Fiber for simpler and faster API handling
- Implemented JWT-based authentication system for user login and protected routes
- Created authentication APIs for user registration and login
- Generated JWT tokens and used them to authorize API requests
- Added middleware to protect task routes using JWT verification
- Integrated Zerolog for structured logging of API requests and responses
- Configured environment variables using `.env` file for DB and JWT secrets
- Tested authentication flow and protected APIs using Postman
- Understood complete request flow including middleware, handlers, and database interaction

</details>

</details>

---

## Tech Used

- Git
- GitLab
- HTML
- CSS
- JavaScript
- Node.js
- Express.js
- Docker
- Go

---

## Notes

This project is part of a learning exercise to understand:


- Repository management
- Web Development
- Integration of front end and backend
- Containerization using Docker
- Backend development using Go