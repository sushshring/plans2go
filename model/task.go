package model

import (
	"fmt"
	"time"
)

// A Task represents a task to be done
type Task struct {
	uid      string
	name     string
	category Category
	due      time.Time
	priority *int64
}

// GetPriority returns the eventual priority value of a task
func (task Task) GetPriority() int64 {
	if task.priority != nil {
		p := *(task.priority)
		return p / time.Until(task.due).Nanoseconds()
	}
	return task.category.priority / time.Until(task.due).Nanoseconds()
}

func (task Task) String() string {
	p := task.category.priority
	if task.priority != nil {
		p = *(task.priority)
	}
	return fmt.Sprintf("{ ID: %s, name: %s, priority: %s, timeDue: %s, category: %s }\n",
		task.uid,
		task.name,
		task.category,
		task.due,
		p,
	)
}
