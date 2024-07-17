package model

type User struct {
	Id       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type UserResult struct {
	Status  string `json:"status"`
	Users   User   `json:"result,omitempty"`
	Message string `json:"message,omitempty"`
}

type UserAllResult struct {
	Status  string  `json:"status"`
	Users   []*User `json:"result,omitempty"`
	Message string  `json:"message,omitempty"`
}

type TokenRes struct {
	Status      string `json:"status"`
	AccessToken string `json:"access_token"`
}
