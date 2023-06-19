package service

import (
	"pankaj-katyare/todo-list/internal/user/model"
	"pankaj-katyare/todo-list/internal/user/repository"
)

func CreateUser(user *model.User) model.UserResult {

	// Save the new user to the database or any other data source
	res := repository.Create(user)

	return model.UserResult{Status: "true", Users: *res}
}

func LoginUser(user *model.User) model.TokenRes {

	res := repository.Login(user)

	return model.TokenRes{Status: "true", AccessToken: res}
}

func GetUser(id string) *model.UserResult {
	res := repository.Get(id)

	return &model.UserResult{Status: "true", Users: res}
}

func GetAllUser() *model.UserAllResult {

	res := repository.GetAll()

	return &model.UserAllResult{Status: "true", Users: res}
}

func UpdateUser(id string, user model.User) model.UserResult {

	res := repository.Update(id, user)

	return model.UserResult{Status: "true", Users: res}
}
