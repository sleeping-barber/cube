package manager

import "github.com/golang-collections/collections/queue"

type Manager struct {
	Pending       queue.Queue
	EventDb       map[string][]TaskEvent
	TaskWorkerMap map[uuid.UUID]string
	TaskDb        map[string][]Task
	WorkerTaskMap map[string][]uuid.UUID
	Workers       []string
}
