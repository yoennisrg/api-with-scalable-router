package task

import (
	"fmt"
)

func findTaskByID(id string) (Task, error) {
	for _, task := range tasks {
		if task.ID == id {
			return task, nil
		}
	}
	return Task{}, fmt.Errorf("%s", "Faild to find task")
}
