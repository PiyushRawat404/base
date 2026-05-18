// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// )

// type User struct {
// 	ID   int    `json:"id"`
// 	Name string `json:"name"`
// 	Age  int    `json:"age"`
// }


// func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {

// 		fmt.Println("Request:", r.Method, r.URL.Path)

// 		next(w, r)
// 	}
// }

// func homeHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "Welcome to server")
// }

// func getUser(w http.ResponseWriter, r *http.Request) {

// 	user := User{
// 		ID:   1,
// 		Name: "Himanshu",
// 		Age:  22,
// 	}

// 	json.NewEncoder(w).Encode(user)
// }

// func createUser(w http.ResponseWriter, r *http.Request) {

// 	var user User

// 	json.NewDecoder(r.Body).Decode(&user)

// 	fmt.Fprintf(w, "User created: %s", user.Name)
// }

// func main() {

// 	http.HandleFunc("/", loggingMiddleware(homeHandler))
// 	http.HandleFunc("/user", loggingMiddleware(getUser))
// 	http.HandleFunc("/create-user", loggingMiddleware(createUser))

// 	fmt.Println("Server running on port 8080")

// 	http.ListenAndServe(":8080", nil)
// }