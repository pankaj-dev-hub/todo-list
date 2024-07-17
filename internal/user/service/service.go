package service

import (
	"fmt"
	"pankaj-katyare/todo-list/internal/user/model"
	"pankaj-katyare/todo-list/internal/user/repository"
)

func CreateUser(user *model.User) model.UserResult {

	// Save the new user to the database or any other data source
	res, err := repository.Create(user)

	status := ""
	if err != nil {
		status = "failed"
	} else {
		status = "true"
	}

	return model.UserResult{Status: status, Users: *res}
}

func LoginUser(user *model.User) model.TokenRes {

	res, err := repository.Login(user)
	if err != nil {
		return model.TokenRes{Status: "false", AccessToken: fmt.Sprintf("%s", err)}
	}

	return model.TokenRes{Status: "true", AccessToken: res}
}

func GetUser(id string) *model.UserResult {
	res, err := repository.Get(id)
	if err != nil {
		return &model.UserResult{Status: "false", Message: fmt.Sprintf("%s", err)}
	}

	return &model.UserResult{Status: "true", Users: *res}
}

func GetAllUser() *model.UserAllResult {

	res, err := repository.GetAll()
	if err != nil {
		return &model.UserAllResult{Status: "true", Message: fmt.Sprintf("%s", err)}
	}

	return &model.UserAllResult{Status: "true", Users: res}
}

func UpdateUser(id string, user model.User) model.UserResult {

	res, err := repository.Update(id, user)
	if err != nil {
		return model.UserResult{Status: "false", Message: fmt.Sprintf("%s", err)}
	}

	return model.UserResult{Status: "true", Users: *res}
}
