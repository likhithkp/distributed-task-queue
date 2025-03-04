package worker

import (
	"context"
	"distributed-task-queue/queue"
	"distributed-task-queue/services"
	"fmt"
	"time"
)

var maxRetries = 3

func FailedTaskQueue() {
	client := queue.Queue()
	retryCount := make(map[string]int)

	for {
		result, err := client.BRPop(context.Background(), time.Second*5, "failed_task_queue").Result()
		if err == nil && len(result) > 0 {
			task := result[1]

			retryCount[task]++

			if retryCount[task] <= maxRetries {
				success := services.ExecuteFailedTask(task)
				if !success {
					client.LPush(context.Background(), "failed_task_queue", task)
					fmt.Println("Retrying task:", task)
				} else {
					delete(retryCount, task)
				}
			} else {
				client.LPush(context.Background(), "discarded_task_queue", task)
				fmt.Println("Task discarded after multiple tries:", task)
				delete(retryCount, task)
			}
		}

		time.Sleep(10 * time.Millisecond)
	}
}
