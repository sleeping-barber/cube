package task

import (
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
	ID uuid.UUID
	Name string
	State State
}
