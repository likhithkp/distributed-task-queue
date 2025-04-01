package producer

import (
	"distributed-task-queue/services"
	"distributed-task-queue/shared"
	"encoding/json"
	"net/http"
)

func Schedule(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		res := &shared.Response{
			Message:    "Not a valid method",
			StatusCode: http.StatusMethodNotAllowed,
		}

		json.NewEncoder(w).Encode(&res)
	}

	task := new(shared.ScheduledTask)

	if err := json.NewDecoder(r.Body).Decode(task); err != nil {
		res := &shared.Response{
			Message:    "Error while decoding data",
			StatusCode: http.StatusInternalServerError,
		}

		json.NewEncoder(w).Encode(&res)
	}

	if task.Name == "" || task.Data == nil {
		res := &shared.Response{
			Message:    "task 'name', 'data' and 'time' are required/missing fields",
			StatusCode: http.StatusMethodNotAllowed,
		}

		json.NewEncoder(w).Encode(&res)
	}

	services.AddScheduledTask(task)
}
