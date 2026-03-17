package main

import (
	"fmt"
	"net/http"
	"task_go/routes"
)

func main() {

	fmt.Println("Server running on port 4000")

	routes.RegisterRoutes()

	http.ListenAndServe(":4000", nil)
}
