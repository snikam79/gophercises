package cmd

import (
	"fmt"
	"os"
	"task/db"

	"github.com/spf13/cobra"
)

// command handler for list command
// not sure how I applied GoF command pattern here ?
var listCommand = &cobra.Command{
	Use:   "list",
	Short: "Lists all tasks from the task list",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.GetAllTasks()
		if err != nil {
			fmt.Println("Something went wrong :", err.Error())
			os.Exit(1)
		}

		if len(tasks) == 0 {
			fmt.Println("No tasks to addess. Take Off !!")
			return
		}

		fmt.Println("You have following tasks to address :")
		for i, task := range tasks {
			fmt.Printf("%d. %s\n", i+1, task.Value)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCommand)
}
