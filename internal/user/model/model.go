package model

type User struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type UserResult struct {
	Status string `json:"status"`
	Users  User   `json:"user"`
}

type UserAllResult struct {
	Status string  `json:"status"`
	Users  []*User `json:"user"`
}
