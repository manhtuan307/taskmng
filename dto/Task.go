package dto

import (
	"time"
)

// A Todo represent task and errand
type Task struct {
	Name      string
	Completed bool
	Created   time.Time
	Updated   time.Time
}
