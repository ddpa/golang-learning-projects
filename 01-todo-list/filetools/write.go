package filetools

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"todo-app/types"
)

func GetReader(filepath string) (*csv.Reader, *os.File, error) {
	f, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to open file for reading")
	}

	r := csv.NewReader(f)

	return r, f, nil
}

func GetWriter(filepath string) (*csv.Writer, *os.File, error) {
	f, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to open file for writing")
	}

	w := csv.NewWriter(f)

	return w, f, nil
}

func AddTask(task types.Task) {
	r, f, err := GetReader("task-list.csv")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	tasksInCsv, err := r.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	f.Close()

	var taskId int

	if len(tasksInCsv) > 1 {
		id := tasksInCsv[len(tasksInCsv)-1][0]
		taskId, err = strconv.Atoi(id)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	} else {
		taskId = 0
	}

	w, fw, err := GetWriter("task-list.csv")
	defer fw.Close()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	w.Write([]string{strconv.Itoa(taskId + 1), task.Name, task.DateCreated, "false"})

	w.Flush()
}
