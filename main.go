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

	http.HandleFunc("/tasks", producer.AddToQueue)
	http.HandleFunc("/tasks/bulk", producer.BulkUpload)
	http.ListenAndServe(":3000", nil)
}
