package model

// A TaskList implements heap.Interface and holds Tasks
type TaskList []*Task

// Len returns the number of tasks in the list
func (taskList TaskList) Len() int {
	return len(taskList)
}

// Less compares two elements in the taskList
func (taskList TaskList) Less(i, j int) bool {
	return taskList[i].getPriority() > taskList[j].getPriority()
}

// Swap swaps two elements in the queue in place
func (taskList TaskList) Swap(i, j int) {
	taskList[i], taskList[j] = taskList[j], taskList[i]
}

// Push adds a new Task to the TaskList
func (taskList *TaskList) Push(x interface{}) {
	item := x.(*Task)
	*taskList = append(*taskList, item)
}

// Pop removes the highest priority item
func (taskList *TaskList) Pop() interface{} {
	old := *taskList
	n := len(*taskList)
	item := old[n-1]
	*taskList = old[0 : n-1]
	return item
}

// Peek returns the top element without removing it
func (taskList *TaskList) Peek() interface{} { return *pq[-1] }

// AsList returns the current elements as a list
func (taskList *TaskList) AsList() []*Task {
	tmp := make([]*Task, len(*taskList))
	copy(*taskList, tmp)
	return tmp
}
