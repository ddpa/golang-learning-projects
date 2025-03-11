/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"todo-app/filetools"
	"todo-app/types"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Task list",
	Long:  `Displays all tasks that need to be completed and all that also have been completed at this point.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list called")
		tasksLoaded, err := filetools.LoadTasks("task-list.csv")

		if err != nil {
			fmt.Println("failed to load records to display tasks")
			os.Exit(1)
		}

		taskList := types.TaskList{
			Tasks: tasksLoaded,
		}

		fmt.Println("teste leitura: ", tasksLoaded[0].Name)

		taskList.PrintTasks()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
