package main

import (
	"distributed-task-queue/producer"
	"distributed-task-queue/queue"
	"net/http"
)

func main() {
	redisClient := queue.Queue()
	defer redisClient.Close()

	http.HandleFunc("/queue", producer.AddToQueue)
	http.ListenAndServe(":3000", nil)
}
