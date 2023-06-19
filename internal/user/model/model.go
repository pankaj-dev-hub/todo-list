package model

type User struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
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

type TokenRes struct {
	Status      string `json:"status"`
	AccessToken string `json:"access_token"`
}
