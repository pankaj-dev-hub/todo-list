package handler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"pankaj-katyare/todo-list/internal/user/model"
	"pankaj-katyare/todo-list/internal/user/service"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {

	var user model.User
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Failed to parse request body =", err)
		json.NewEncoder(w).Encode(map[string]bool{"Status": false})
	}
	err = json.Unmarshal(data, &user)
	if err != nil {
		_ = json.NewEncoder(w).Encode(map[string]bool{"status": false})
	}

	log.Println("Request data: ", user)

	// Service
	res := service.CreateUser(&user)

	log.Println("Response from the service: ", res)

	// Return a success response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)

}

func GetUser(w http.ResponseWriter, r *http.Request) {

	id := r.FormValue("id")

	if id == "" {
		log.Fatal("Id cannot be empty...")
	}

	log.Println("task id:", id)
	res := service.GetUser(id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func GetAllUser(w http.ResponseWriter, r *http.Request) {

	res := service.GetAllUser()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

	var user model.User

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Failed to parse request body =", err)
		json.NewEncoder(w).Encode(map[string]bool{"Status": false})
	}
	err = json.Unmarshal(data, &user)
	if err != nil {
		log.Println("Failed to parse request body =", err)
		json.NewEncoder(w).Encode(map[string]bool{"Status": false})
	}

	id := r.FormValue("id")

	if id == "" {
		log.Fatal("Id cannot be empty...")
	}

	res := service.UpdateUser(id, user)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
