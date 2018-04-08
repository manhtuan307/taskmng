package dto

type GetTasksActionResult struct {
	IsSucess   bool
	Message    string
	Tasks      []Task
	NumOfPages int
	PageIndex  int
}
