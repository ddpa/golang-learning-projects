/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"time"

	"todo-app/types"

	"github.com/spf13/cobra"

	"todo-app/filetools"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add Task",
	Long:  `A longer description that`,
	Run: func(cmd *cobra.Command, args []string) {
		taskName, _ := cmd.Flags().GetString("task")

		task := types.Task{
			Name:        taskName,
			DateCreated: time.Now().UTC().Format(time.RFC3339),
		}

		fmt.Println("Adding Task - ", task.Name)

		filetools.AddTask(task)

	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:

	addCmd.PersistentFlags().String("task", "", "task name to be added")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
