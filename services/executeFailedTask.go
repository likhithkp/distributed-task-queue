package services

import (
	"fmt"
)

func ExecuteFailedTask(task string) bool {
	go fmt.Println("FAILED TASK EXECUTED", task)
	return true
}
