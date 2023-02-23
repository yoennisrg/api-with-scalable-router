package task

type Task struct {
	ID          string `json:"id"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Completed   bool   `json:"completed"`
}

var tasks = []Task{
	{ID: "1", Title: "Task 1", Description: "This is task 1", Completed: false},
	{ID: "2", Title: "Task 2", Description: "This is task 2", Completed: false},
	{ID: "3", Title: "Task 3", Description: "This is task 3", Completed: true},
}
