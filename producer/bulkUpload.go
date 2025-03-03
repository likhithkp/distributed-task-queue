package producer

import (
	"distributed-task-queue/services"
	"distributed-task-queue/shared"
	"encoding/json"
	"net/http"
)

func BulkUpload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		res := &shared.Response{
			Message:    "Not a valid method",
			StatusCode: http.StatusMethodNotAllowed,
		}

		json.NewEncoder(w).Encode(&res)
	}

	tasks := new(shared.BulkTask)

	if err := json.NewDecoder(r.Body).Decode(&tasks); err != nil {
		res := &shared.Response{
			Message:    "Error while decoding data",
			StatusCode: http.StatusInternalServerError,
		}

		json.NewEncoder(w).Encode(&res)
	}

	if tasks.Name == "" || len(tasks.Tasks) == 0 {
		res := &shared.Response{
			Message:    "task 'name' and 'data' are required/missing fields",
			StatusCode: http.StatusMethodNotAllowed,
		}

		json.NewEncoder(w).Encode(&res)
	}

	for _, task := range tasks.Tasks {
		if task.Priority {
			services.AddToHighPriorityQueue(&task)
		} else {
			services.AddToLowPriorityQueue(&task)
		}
	}
}
