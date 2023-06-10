package model

type Task struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
	Status      string `json:"status"`
}

type Todo struct {
	Tasks []Task `json:"tasks"`
}

type TodoResult struct {
	Status string `json:"status"`
	Todo   Todo   `json:"todo"`
}

type TodoAllResult struct {
	Status string  `json:"status"`
	Todo   []*Todo `json:"todo"`
}
