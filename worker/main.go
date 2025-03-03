package worker

import (
	"context"
	"distributed-task-queue/queue"
	"distributed-task-queue/services"
	"fmt"
	"log"
	"time"
)

func Worker() {
	client := queue.Queue()

	for {
		highPriorityQueueLength, err := client.LLen(context.Background(), "high_priority_queue").Result()
		if err != nil {
			log.Println("Error while getting the length the high_priority_queue")
			return
		}

		normalPriorityQueueLength, err := client.LLen(context.Background(), "normal_priority_queue").Result()
		if err != nil {
			log.Println("Error while getting the length the normal_priority_queue")
			return
		}

		if highPriorityQueueLength > 0 {
			task, err := client.BRPop(context.Background(), time.Second*5, "high_priority_queue").Result()

			if err != nil {
				log.Println("Error while popping the task from high_priority_queue")
				return
			}

			services.ExecuteTask(&task)
		} else if highPriorityQueueLength == 0 && normalPriorityQueueLength == 0 {
			fmt.Println("WAITING FOR TASK")
		} else {
			if normalPriorityQueueLength > 0 {
				normalTask, err := client.BRPop(context.Background(), time.Second*5, "normal_priority_queue").Result()

				if err != nil {
					log.Println("Error while popping the task from normal_priority_queue")
					return
				}

				services.ExecuteTask(&normalTask)
			}
		}
		fmt.Println("HIGH PRIORITY QUEUE LENGTH", highPriorityQueueLength)
		fmt.Println("NORMAL PRIORITY QUEUE LENGTH", normalPriorityQueueLength)
		// time.Sleep(10 * time.Millisecond)
		time.Sleep(time.Second * 1)
	}
}
