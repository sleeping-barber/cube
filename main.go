package main

import (
	"fmt"
	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
	"github.com/sleeping-barber/cube/manager"
	"github.com/sleeping-barber/cube/node"
	"github.com/sleeping-barber/cube/task"
	"github.com/sleeping-barber/cube/worker"
	"time"
)

func main() {
	// create Task object
	t := task.Task{
		ID:            uuid.New(),
		Name:          "Task",
		State:         task.Pending,
		Image:         "Image",
		Memory:        1024,
		Disk:          1,
		ExposedPort:   nil,
		PortBindings:  nil,
		RestartPolicy: "",
		StartTime:     time.Now(),
		FinishTime:    time.Now(),
	}

	// create TaskEvent object
	te := task.TaskEvent{
		ID:        uuid.New(),
		State:     task.Pending,
		Timestamp: time.Now(),
		Task:      t,
	}

	// print Task and TaskEvent objects
	fmt.Printf("task: %v\n", t)
	fmt.Printf("task event: %v\n", te)
	fmt.Println()

	// create Worker object
	w := worker.Worker{
		Name:      "worker",
		Queue:     *queue.New(),
		Db:        make(map[uuid.UUID]task.Task),
		TaskCount: 0,
	}

	// print Worker object
	fmt.Printf("worker: %v\n", w)
	fmt.Println()

	// run Worker methods
	w.CollectStats()
	w.RunTask()
	w.StartTask()
	w.StopTask()

	// create Manager object
	m := manager.Manager{
		Pending:       *queue.New(),
		EventDb:       make(map[string][]task.TaskEvent),
		TaskWorkerMap: nil,
		TaskDb:        make(map[string][]task.Task),
		WorkerTaskMap: nil,
		Workers:       []string{w.Name},
	}

	// print Manager object
	fmt.Printf("manager: %v\n", m)
	// call Manager methods
	m.SelectWorker()
	m.UpdateTask()
	m.SendWork()
	fmt.Println()

	// create Node object
	n := node.Node{
		Name:            "node",
		Ip:              "192.168.1.1",
		Memory:          1024,
		Disk:            25,
		MemoryAllocated: 0,
		DiskAllocated:   0,
		TasksCount:      0,
	}

	// print Node object
	fmt.Println(n)
}
