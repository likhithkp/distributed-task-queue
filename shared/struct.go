package shared

type Response struct {
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
}

type BulkTask struct {
	Name  string
	Tasks []Task
}

type Task struct {
	Name     string
	Data     any
	Priority bool
}
