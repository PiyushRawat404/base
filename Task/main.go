package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

const fileName = "tasks.json"

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter task: ")
	input, _ := reader.ReadString('\n')

	title := strings.TrimSpace(input)

	addTask(title)
}

func addTask(title string) {
	tasks := loadTasks()

	newTask := Task{
		ID:    len(tasks) + 1,
		Title: title,
	}

	tasks = append(tasks, newTask)
	saveTasks(tasks)

	fmt.Println("Task added!")
}

func loadTasks() []Task {
	var tasks []Task

	data, err := os.ReadFile(fileName)
	if err != nil {
		return tasks
	}

	json.Unmarshal(data, &tasks)
	return tasks
}

func saveTasks(tasks []Task) {
	data, _ := json.MarshalIndent(tasks, "", "  ")
	os.WriteFile(fileName, data, 0644)
}
