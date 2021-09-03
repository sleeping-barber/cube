package task

import (
	"time"

	"github.com/docker/go-connections/nat"
	"github.com/google/uuid"
)

type State int

const (
	Completed State = iota
	Failed
	Pending
	Running
	Scheduled
)

type Task struct {
	ID            uuid.UUID
	Name          string
	State         State
	Image         string
	Memory        int
	Disk          int
	ExposedPort   nat.PortSet
	PortBindings  map[string]string
	RestartPolicy string
	StartTime     time.Time
	FinishTime    time.Time
}

type TaskEvent struct {
	ID        uuid.UUID
	State     State
	Timestamp time.Time
	Task      Task
}
