package model

import (
	"fmt"
)

// A Category represents categories of tasks that are possible.
// Each category has a certain priority value that is set
// by the user
type Category struct {
	uid      string
	name     string
	priority int64
}

func (category Category) String() string {
	return fmt.Sprintf("{ ID: %s, name: %s, priority: %d}\n", category.uid, category.name, category.priority)
}
