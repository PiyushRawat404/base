package http


import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}


func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to  server")
}
func getUser(w http.ResponseWriter, r *http.Request) {

	user := User{
		ID:   1,
		Name: "Himanshu",
		Age:  22,
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(user)
}


func createUser(w http.ResponseWriter, r *http.Request) {

	var user User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}


	if user.Name == "" {
		http.Error(w, "Name is required", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "User created successfully: %s", user.Name)
}

func main() {

	http.HandleFunc("/", homeHandler)


	http.HandleFunc("/user", getUser)

	http.HandleFunc("/create-user", createUser)

	fmt.Println("Server running on port 8080")

	http.ListenAndServe(":8080", nil)
}