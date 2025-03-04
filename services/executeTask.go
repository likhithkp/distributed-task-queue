package services

import (
	"fmt"
)

func ExecuteTask(task string) bool {
	go fmt.Println("TASK EXECUTED", task)
	return true
}
