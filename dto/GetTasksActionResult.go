package dto

type GetTasksActionResult struct {
	Result     *ActionResult
	Tasks      []Task
	NumOfPages int
	PageIndex  int
}
