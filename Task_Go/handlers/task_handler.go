package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"task_go/models"
	"task_go/storage"
)

func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {

	var task models.Task

	json.NewDecoder(r.Body).Decode(&task)

	task.ID = len(storage.Tasks) + 1

	storage.Tasks = append(storage.Tasks, task)

	json.NewEncoder(w).Encode(task)
	fmt.Println(storage.Tasks)

}

func GetTasksHandler(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(storage.Tasks)
	fmt.Println(storage.Tasks)

}
