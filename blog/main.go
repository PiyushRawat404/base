package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

type Blog struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var db *sql.DB

func createBlog(w http.ResponseWriter, r *http.Request) {

	var blog Blog

	json.NewDecoder(r.Body).Decode(&blog)

	query := `
		INSERT INTO posts(title, content)
		VALUES($1, $2)
	`

	_, err := db.Exec(query, blog.Title, blog.Content)

	if err != nil {
		http.Error(w, "Failed to create blog", http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "Blog created successfully")
}

func main() {

	connStr := "user=himanshu password=1234 dbname=blogdb sslmode=disable"

	var err error

	db, err = sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to PostgreSQL successfully!")

	http.HandleFunc("/create-blog", createBlog)

	fmt.Println("Server running on port 8080")

	http.ListenAndServe(":8080", nil)
}
