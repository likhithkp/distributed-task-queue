package shared

type Response struct {
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
}

type BulkTask struct {
	Name     string `json:"name"`
	Function func() `json:"function"`
	Tasks    []Task `json:"tasks"`
}

type Task struct {
	Name     string `json:"name"`
	Data     any    `json:"data"`
	Priority bool   `json:"priority"`
}
