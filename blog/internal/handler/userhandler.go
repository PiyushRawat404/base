package handler

import (
	"blog/internal/model"
	"encoding/json"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request){
	if r.Method!=http.MethodPost{
		writeJSON(w,http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}

	var user model.User

	if err :=json.NewDecoder(r.Body).Decode(&user)
	err !=nil{
		writeJSON(w,http.StatusBadRequest,map[string]string{"error":"invalid body"})
		return
	}

	
	writeJSON(w,http.StatusCreated,map[string]string{"sucess":"user created"})
}