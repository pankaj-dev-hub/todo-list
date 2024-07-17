package service

import (
	"fmt"
	"log"
	"pankaj-katyare/todo-list/internal/todo/model"
	"pankaj-katyare/todo-list/internal/todo/repository"
)

func CreateTask(todo *model.Todo) model.TodoResult {

	// Save the new user to the database or any other data source
	res, err := repository.Create(todo)
	if err != nil {
		return model.TodoResult{Status: "false", Message: fmt.Sprintf("%s", err)}
	}

	return model.TodoResult{Status: "true", Todo: *res}
}

func GetTask(id string) *model.TodoResult {

	res, err := repository.Get(id)
	if err != nil {
		return &model.TodoResult{Status: "false", Message: fmt.Sprintf("%s", err)}
	}

	var result model.Todo
	log.Println("lenght of task =", len(res.Tasks))
	for r, p := range res.Tasks {
		log.Printf("\n %d = %+v :", r, p)
		if id == res.Tasks[r].Id {
			result.Tasks = append(result.Tasks, p)
		}
	}

	return &model.TodoResult{Status: "true", Todo: result}

}

func GetAllTask() *model.TodoAllResult {

	res, err := repository.GetAll()
	if err != nil {
		return &model.TodoAllResult{Status: "false", Message: fmt.Sprintf("%s", err)}
	}

	return &model.TodoAllResult{Status: "true", Todo: res}
}

func UpdateTask(id string, todo model.Todo) model.TodoResult {

	res, err := repository.Update(id, todo)
	if err != nil {
		return model.TodoResult{Status: "false", Message: fmt.Sprintf("%s", err)}
	}

	return model.TodoResult{Status: "true", Todo: *res}
}

func CompletedTask() *model.TodoAllResult {

	var todos []*model.Todo

	res, err := repository.GetCompletedTask()
	if err != nil {
		return &model.TodoAllResult{Status: "false", Message: fmt.Sprintf("%s", err)}
	}

	flag := 0
	for _, todo := range res {
		var tasks []model.Task
		flag = 0
		for _, task := range todo.Tasks {
			if task.Status == "completed" {
				tasks = append(tasks, task)
				flag = 1
				log.Println("Task :", task)
			}
		}
		if flag == 1 {
			todo = &model.Todo{
				Tasks: tasks,
			}
			todos = append(todos, todo)
			log.Println("Todo :", todos)
		}
	}

	return &model.TodoAllResult{Status: "true", Todo: todos}
}

func PendingTask() *model.TodoAllResult {

	var todos []*model.Todo

	res, err := repository.GetPendingTask()
	if err != nil {
		return &model.TodoAllResult{Status: "false", Message: fmt.Sprintf("%s", err)}
	}

	flag := 0
	for _, todo := range res {
		var tasks []model.Task
		flag = 0
		for _, task := range todo.Tasks {
			if task.Status == "pending" {
				tasks = append(tasks, task)
				flag = 1
				log.Println("Task :", task)
			}
		}
		if flag == 1 {
			todo = &model.Todo{
				Tasks: tasks,
			}
			todos = append(todos, todo)
			log.Println("Todo :", todos)
		}
	}

	return &model.TodoAllResult{Status: "true", Todo: todos}
}
