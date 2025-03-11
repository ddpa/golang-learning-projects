package types

import (
	"fmt"
	"os"
	"text/tabwriter"
)

//type Tasks []Task

type Task struct {
	Id          int
	Name        string
	DateCreated string
	Done        bool
}

type TaskList struct {
	Tasks []Task
}

type TaskPrint interface {
	PrintTasks()
}

func (t *TaskList) PrintTasks() {
	fmt.Println("display function")
	w := tabwriter.NewWriter(os.Stdout, 15, 20, 10, ' ', 0)
	defer w.Flush()

	fmt.Fprintln(w, "----\t----\t----\t----")
	fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", "ID", "NAME", "DATE", "DONE")
	fmt.Fprintln(w, "----\t----\t----\t----")

	for _, task := range t.Tasks {
		fmt.Fprintf(w, "%d\t%s\t%s\t%t\n", task.Id, task.Name, task.DateCreated, task.Done)
	}
}
