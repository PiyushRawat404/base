package routes

import (
	"net/http"
	"task_go/handlers"
)

func RegisterRoutes() {

	http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "GET" {
			handlers.GetTasksHandler(w, r)
		}

		if r.Method == "POST" {
			handlers.CreateTaskHandler(w, r)
		}

	})
}