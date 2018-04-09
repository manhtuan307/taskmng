package dto

// A GetTasksActionResult represent list of tasks together with paging info, etc
type GetTasksActionResult struct {
	IsSuccess  bool
	Message    string
	Tasks      []Task
	NumOfPages int
	PageIndex  int
}
