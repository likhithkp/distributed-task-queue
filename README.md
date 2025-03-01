# Distributed Task Queue

## Overview  
Distributed Task Queue is a lightweight, high-performance distributed task queue using **Go, Redis, and net/http**. It allows you to queue tasks and process them asynchronously with worker nodes.

## Features  
✅ **Fast & Lightweight** – Uses Redis Lists for queueing.  
✅ **FIFO Processing** – Tasks are executed in the order they arrive.  
✅ **Scalable** – Supports multiple producers and workers.  
✅ **Fault-Tolerant** – Ensures reliable processing with Redis persistence.  

## Architecture  
- **Producer (API Endpoint)** – Accepts tasks via an HTTP request and pushes them to Redis.  
- **Redis (Queue)** – Stores tasks in a **list** using `LPUSH` and `BRPOP`.  
- **Worker (Task Executor)** – Continuously polls Redis, pulls tasks, and processes them.  

## Installation  
1. Clone the repo:  
   ```sh
   git clone https://github.com/likhithkp/distributed-task-queue.git  
   cd redis-task-queue  
   ```  
2. Install dependencies:  
   ```sh
   go mod tidy  
   ```  
3. Start Redis:  
   ```sh
   docker run -p 6379:6379 redis  
   ```  
4. Run the producer:  
   ```sh
   go run producer/main.go  
   ```  
5. Run the worker:  
   ```sh
   go run worker/main.go  
   ```  

## Usage  
### **1️⃣ Add a Task (Producer API)**  
Send a task to the queue using cURL or Postman:  
```sh
curl -X POST "http://localhost:3000/enqueue" -d '{"task": "send_email:user123"}'
```  

### **2️⃣ Process Tasks (Worker)**  
The worker automatically pulls tasks from Redis and executes them.  

## Configuration  
| Environment Variable | Description | Default |  
|----------------------|-------------|---------|  
| `REDIS_URL` | Redis connection string | `localhost:6379` |  
| `QUEUE_NAME` | Name of the Redis list | `task_queue` |  

## License  
🛠️ MIT License – Feel free to use and modify.  
