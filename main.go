package main

import (
	"distributed-task-queue/producer"
	"distributed-task-queue/queue"
	"distributed-task-queue/worker"
	"net/http"
)

func main() {
	redisClient := queue.Queue()
	defer redisClient.Close()

	go worker.Worker()
	go worker.ScheduledTask()
	go worker.FailedTaskQueue()

	http.HandleFunc("/tasks", producer.AddToQueue)
	http.HandleFunc("/tasks/bulk", producer.BulkUpload)
	http.HandleFunc("/tasks/schedule", producer.Schedule)
	http.ListenAndServe(":3000", nil)
}
