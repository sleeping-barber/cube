package main

import (
	"fmt"
	"github.com/docker/docker/client"
	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
	"github.com/sleeping-barber/cube/manager"
	"github.com/sleeping-barber/cube/node"
	"github.com/sleeping-barber/cube/task"
	"github.com/sleeping-barber/cube/worker"
	"os"
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

	fmt.Printf("Create a test container\n")
	dockerTask, createResult := createContainer()
	if createResult.Error != nil {
		fmt.Printf("%v\n", createResult.Error)
		os.Exit(1)
	}

	time.Sleep(time.Second * 5)
	fmt.Printf("stopping container %s\n", createResult.ContainerId)
	_ = stopContainer(dockerTask)
}

func createContainer() (*task.Docker, *task.DockerResult) {
	c := task.Config{
		Name:  "test-container-1",
		Image: "postgres:13",
		Env: []string{
			"POSTGRES_USER=cube",
			"POSTGRES_PASSWORD=secret",
		},
	}

	dc, _ := client.NewClientWithOpts(client.FromEnv)

	d := task.Docker{
		Config: c,
		Client: dc,
	}

	result := d.Run()
	if result.Error != nil {
		fmt.Printf("%v\n", result.Error)
		return nil, nil
	}

	fmt.Printf("Container %s is running with config %v\n", result.ContainerId, c)
	return &d, &result
}

func stopContainer(d *task.Docker) *task.DockerResult {
	result := d.Stop()
	if result.Error != nil {
		fmt.Printf("%v\n", result.Error)
		return nil
	}
	fmt.Printf("Container %s has been stopped and removed\n", d.ContainerId)
	return &result
}
