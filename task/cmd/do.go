package cmd

import (
	"fmt"
	"strconv"
	"task/db"

	"github.com/spf13/cobra"
)

// go file for do command
// it simply marks the task as completed by removing task from bolt db

var doCommand = &cobra.Command{
	Use:   "do",
	Short: "Performs task in task list and marks it as complete",
	// closure usage
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int

		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Unable to parse the arrument : ", arg)
			} else {
				ids = append(ids, id)
			}
		}
		// now get all the tasks from bolt db
		tasks, err := db.GetAllTasks()
		if err != nil {
			fmt.Println("Something went wrong", err)
			return
		}

		for _, id := range ids {
			if id <= 0 || id > len(tasks) {
				fmt.Println("Invalid task number :", id)
				continue
			}
			task := tasks[id-1]
			err := db.DeleteTask(task.Key)
			if err != nil {
				fmt.Printf("Unable to mark \"%d\" as complete. Error :%s\n", id, err)
			} else {
				fmt.Printf("Marked \"%d\" as complete", id)
			}

		}
	},
}

func init() {
	RootCmd.AddCommand(doCommand)
}
