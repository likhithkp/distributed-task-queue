package services

import (
	"context"
	"distributed-task-queue/queue"
	"distributed-task-queue/shared"
	"encoding/json"
	"log"
)

func AddToHighPriorityQueue(task *shared.Task) {
	byteData, err := json.Marshal(task)

	if err != nil {
		log.Println("Error while serializing task", err.Error())
		return
	}

	queue.Queue().LPush(context.Background(), "high_priority_queue", string(byteData))
}
