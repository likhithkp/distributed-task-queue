
# Distributed Task Queue

A simple distributed task queue implementation in Go using Redis, where tasks are added, processed, and prioritized for execution by workers. This system demonstrates how to use Redis lists to manage task queues, enabling both high and low priority tasks to be processed efficiently.

## Features
- **Task Queueing**: Tasks can be added to a queue with either high or normal priority.
- **Task Scheduling**: Tasks can be scheduled to a queue.
- **Worker**: A worker processes tasks based on their priority (high priority tasks first).
- **REST API**: A simple HTTP interface for adding tasks to the queue and bulk uploading tasks.

## Technologies
- **Go**: The programming language for implementing the producer, worker, and queue systems.
- **Redis**: Used as the underlying message queue to manage and store tasks.
- **HTTP Server**: Simple HTTP handlers to interact with the task queue through API endpoints.

## Requirements
- Go version 1.18 or later
- Redis (installed locally or accessible via a URL)

## Setup and Installation

### Step 1: Clone the Repository
```bash
git clone https://github.com/likhithkp/distributed-task-queue.git
cd distributed-task-queue
```

### Step 2: Install Dependencies
Ensure that you have Redis running. If Redis is not installed locally, you can either [download and install Redis](https://redis.io/download) or use a hosted solution.

Install required Go packages:
```bash
go mod tidy
```

### Step 3: Run the Application
You can start the application by running:

```bash
go run main.go
```

This will start the HTTP server on `http://localhost:3000`.

### Step 4: Access the API

#### Add a Task to the Queue
You can add a task to the queue by sending a POST request to `/tasks`. For example, using `curl`:

```bash
curl -X POST http://localhost:3000/tasks -d '{"priority":true,"task":"high-priority task"}'
```

This will add a task with `high-priority` to the Redis queue.

#### Bulk Upload Tasks
You can also upload multiple tasks at once using the `/tasks/bulk` endpoint:

```bash
curl -X POST http://localhost:3000/tasks/bulk -d '[{"priority":true,"task":"task 1"}, {"priority":false,"task":"task 2"}]'
```

This will add multiple tasks in one go, some of them with high priority, others with normal priority.

### Step 5: Worker Logic
The worker will continuously listen for tasks in the queue and process them. The worker prioritizes high-priority tasks and processes them first. After all high-priority tasks are processed, it will begin processing normal-priority tasks.

## Code Structure
- **`main.go`**: Entry point of the application, setting up the Redis connection, starting the worker, and initializing the HTTP server.
- **`producer`**: Contains the functions responsible for adding tasks to the queue (`AddToQueue` and `BulkUpload`).
- **`queue`**: Manages the Redis client connection and queue operations.
- **`worker`**: Contains the worker logic for fetching and processing tasks from the queue.
- **`shared`**: Contains common structs, such as `Task`.

## API Endpoints

### `POST /tasks`
- **Description**: Adds a single task to the queue.
- **Request Body**:
  ```json
  {
    "priority": true,
    "task": "Task description"
  }
  ```
- **Response**: HTTP status `200 OK` if the task was added successfully.

### `POST /tasks/bulk`
- **Description**: Adds multiple tasks to the queue in bulk.
- **Request Body**:
  ```json
  [
    {
      "priority": true,
      "task": "Task 1"
    },
    {
      "priority": false,
      "task": "Task 2"
    }
  ]
  ```
- **Response**: HTTP status `200 OK` if all tasks were added successfully.

### `POST /tasks/schedule`
- **Description**: Schedule task.
- **Request Body**:
  ```json
  {
    "priority": true,
    "task": "Task description",
    "time": "2006-01-02T15:04:05Z07:00"
  }
  ```
- **Response**: HTTP status `200 OK` if the task was added successfully.

## Worker Behavior
- The worker processes tasks from the Redis queue:
  - First, it processes high-priority tasks.
  - Once all high-priority tasks are completed, it moves to process normal-priority tasks.
- Tasks are processed by popping them from the queue and executing them. If a task fails, it will be moved to a dead-letter queue (DLQ) for retries.
- Scheduled tasks are also handled by checking the scheduled queue periodically and executing tasks when their scheduled time arrives.

## Scaling and Improvements
- **Multiple Workers**: You can scale this by deploying multiple workers to handle more tasks concurrently.
- **Retry Logic**: Implementing retries for failed tasks or failed processing attempts.
- **Metrics & Monitoring**: Add logging and monitoring for better visibility of task processing.

## Conclusion
This distributed task queue can be a base system for task scheduling in applications where processing order and prioritization are key.
