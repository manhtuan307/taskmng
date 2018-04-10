package dto

// A SearchTasksActionResult represent list of tasks together with paging info, etc
type SearchTasksActionResult struct {
	IsSuccess  bool
	Message    string
	Tasks      []Task
	NumOfPages int
	PageIndex  int
}
