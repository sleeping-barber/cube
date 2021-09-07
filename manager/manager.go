package manager

import (
	"fmt"
	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
	"github.com/sleeping-barber/cube/task"
)

type Manager struct {
	Pending       queue.Queue
	EventDb       map[string][]task.TaskEvent
	TaskWorkerMap map[uuid.UUID]string
	TaskDb        map[string][]task.Task
	WorkerTaskMap map[string][]uuid.UUID
	Workers       []string
}

func (m *Manager) SelectWorker() {
	fmt.Println("I will select an appropriate worker")
}

func (m *Manager) UpdateTask() {
	fmt.Println("I will update tasks")
}

func (m *Manager) SendWork() {
	fmt.Println("I will send work to workers")
}
