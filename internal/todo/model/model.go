package model

type Task struct {
	Id          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	StartDate   string `json:"start_date,omitempty"`
	EndDate     string `json:"end_date,omitempty"`
	Status      string `json:"status,omitempty"`
}

type Todo struct {
	Tasks []Task `json:"tasks,omitempty"`
}

type TodoResult struct {
	Status  string `json:"status"`
	Todo    Todo   `json:"result,omitempty"`
	Message string `json:"message,omitempty"`
}

type TodoAllResult struct {
	Status  string  `json:"status"`
	Todo    []*Todo `json:"result,omitempty"`
	Message string  `json:"message,omitempty"`
}
