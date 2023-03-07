package cli

import (
	"fmt"
	"task/types"
)

func printTasks(tasks []*types.Task) {
	for _, task := range tasks {
		fmt.Printf("%d - %s\n", task.Id, task.Text)
	}
}
