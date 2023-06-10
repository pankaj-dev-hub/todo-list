package handler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"pankaj-katyare/todo-list/internal/todo/model"
	"pankaj-katyare/todo-list/internal/todo/service"
)

func CreateTask(w http.ResponseWriter, r *http.Request) {

	var todo model.Todo
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Failed to parse request body =", err)
		json.NewEncoder(w).Encode(map[string]bool{"Status": false})
	}
	err = json.Unmarshal(data, &todo)
	if err != nil {
		_ = json.NewEncoder(w).Encode(map[string]bool{"status": false})
	}

	log.Println("Request data: ", todo)

	// Service
	res := service.CreateTask(&todo)

	log.Println("Response from the service: ", res)

	// Return a success response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)

}

func GetTask(w http.ResponseWriter, r *http.Request) {

	id := r.FormValue("id")

	if id == "" {
		log.Fatal("Id cannot be empty...")
	}

	log.Println("task id:", id)
	res := service.GetTask(id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func GetAllTask(w http.ResponseWriter, r *http.Request) {

	res := service.GetAllTask()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {

	var todo model.Todo

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Failed to parse request body =", err)
		json.NewEncoder(w).Encode(map[string]bool{"Status": false})
	}
	err = json.Unmarshal(data, &todo)
	if err != nil {
		log.Println("Failed to parse request body =", err)
		json.NewEncoder(w).Encode(map[string]bool{"Status": false})
	}

	id := r.FormValue("id")

	if id == "" {
		log.Fatal("Id cannot be empty...")
	}

	res := service.UpdateTask(id, todo)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func CompletedTask(w http.ResponseWriter, r *http.Request) {

	res := service.CompletedTask()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func PendingTask(w http.ResponseWriter, r *http.Request) {

	res := service.PendingTask()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
