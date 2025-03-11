package filetools

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"todo-app/types"
)

func LoadTasks(filepath string) ([]types.Task, error) {

	var tasks []types.Task

	f, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("failed to open file for reading")
	}

	r := csv.NewReader(f)

	defer f.Close()

	taskRecords, err := r.ReadAll()
	if err != nil {
		return tasks, fmt.Errorf("failed to read records from csv: %w", err)
	}

	//header := taskRecords[0]

	for i, task := range taskRecords {
		if i == 0 {
			continue
		}
		if len(task) != 4 {
			return tasks, fmt.Errorf("invalid task format")
		}

		doneValue := task[3]
		var done bool
		var id int

		id, err = strconv.Atoi(task[0])
		if err != nil {
			return tasks, fmt.Errorf("failed to convert Id to integer")
		}

		if doneValue == "true" || doneValue == "false" {
			donevalue, err := strconv.ParseBool(doneValue)
			if err != nil {
				return tasks, fmt.Errorf("failed to convert the Done column value")
			}
			done = donevalue
		} else {
			return tasks, fmt.Errorf("invalid Done column value")
		}

		tasks = append(tasks, types.Task{
			Id:          id,
			Name:        task[1],
			DateCreated: task[2],
			Done:        done,
		})

	}

	return tasks, nil

}
