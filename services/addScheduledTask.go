package services

import (
	"context"
	"distributed-task-queue/queue"
	"distributed-task-queue/shared"
	"encoding/json"
	"log"
)

func AddScheduledTask(task *shared.ScheduledTask) {
	byteData, err := json.Marshal(task)

	if err != nil {
		log.Println("Error while serializing task", err.Error())
		return
	}

	queue.Queue().LPush(context.Background(), "scheduled_task_queue", string(byteData))
}
