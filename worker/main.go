package worker

import (
	"context"
	"distributed-task-queue/queue"
	"distributed-task-queue/services"
	"fmt"
	"time"
)

func Worker() {
	client := queue.Queue()

	for {
		result, err := client.BRPop(context.Background(), time.Second*5, "high_priority_queue", "normal_priority_queue").Result()
		if err == nil && len(result) > 0 {
			task := result[1]
			success := services.ExecuteTask(task)
			if !success {
				client.LPush(context.Background(), "failed_task_queue", task)
			}
		} else {
			fmt.Println("Worker waiting...")
		}

		time.Sleep(10 * time.Millisecond)
	}
}
